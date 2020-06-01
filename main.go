package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"os"

	"github.com/sensu/sensu-go/types"
	"github.com/sfreiberg/gotwilio"
	"github.com/spf13/cobra"
)

//Declare our variable types here
var (
	accountSid string
	authToken  string
	fromNumber string
	recipients  string
	stdin      *os.File
)

//Start main function
func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

//Configure the root command and add our flags
func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-go-twilio-handler",
		Short: "The Sensu Go Handler for Twilio",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&accountSid,
		"accountSid",
		"s",
		os.Getenv("TWILIO_ACCOUNT_SID"),
		"The account SID for your Twilio account, uses the environment variable TWILIO_ACCOUNT_SID by default")

	cmd.Flags().StringVarP(&authToken,
		"authToken",
		"t",
		os.Getenv("TWILIO_AUTH_TOKEN"),
		"The authorization token for your Twilio account, uses the environment variable TWILIO_AUTH_TOKEN by default")

	cmd.Flags().StringVarP(&fromNumber,
		"fromNumber",
		"f",
		"",
		"Your Twilio phone number")

	_ = cmd.MarkFlagRequired("fromNumber")

	cmd.Flags().StringVarP(&recipients,
		"recipients",
		"r",
		"",
		"The recipient phone number(s), separated by comma")

	_ = cmd.MarkFlagRequired("recipients")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}

	eventJSON, err := ioutil.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %s", err)
	}

	event := &types.Event{}
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal stdin data: %s", err)
	}

	if err = event.Validate(); err != nil {
		return fmt.Errorf("failed to validate event: %s", err)
	}

	if !event.HasCheck() {
		return fmt.Errorf("event does not contain check")
	}

	return sendSMS(event)
}

//This function will send an SMS when receive an alert in error state.
func sendSMS(event *types.Event) error {

	//Set up our message we want to send
	message := "Sensu alert for " + event.Check.Name + " on " + event.Entity.Name + ". Check output: " + event.Check.Output

	//Set up a Twilio client with our accountSid & authToken
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	recipients := strings.Split(recipients, ",")
	for i := range recipients {
		//Send our message to our recipients
		twilio.SendSMS(fromNumber, recipients[i], message, "", "")
	}

	return nil
}
