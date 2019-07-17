package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
//	"github.com/sfreiberg/gotwilio"
)

//Declare our variable types here
var (
	accountSid string
	authToken  string
	fromNumber string
	recipient  string
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

	_ = cmd.MarkFlagRequired("accountSid")

	cmd.Flags().StringVarP(&authToken,
		"authToken",
		"t",
		os.Getenv("TWILIO_AUTH_TOKEN"),
		"The authorization token for your Twilio account, uses the environment variable TWILIO_AUTH_TOKEN by default")

	_ = cmd.MarkFlagRequired("authToken")

	cmd.Flags().StringVarP(&fromNumber,
		"fromNumber",
		"f",
		"",
		"Your Twilio phone number")

	_ = cmd.MarkFlagRequired("fromNumber")

	cmd.Flags().StringVarP(&recipient,
		"recipient",
		"r",
		"",
		"The recipient's phone number")

	_ = cmd.MarkFlagRequired("recipient")

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

	return sendText(event)
}

func sendText(event *types.Event) error {

	//Set our API URL
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	//Set up our message we want to send
	msg := "Sensu alert for " + event.Check.Name + " on " + event.Entity.Name + ". Check output: " +  event.Check.Output

	//Set up our message data
	msgData := url.Values{}
	msgData.Set("To",recipient)
	msgData.Set("From",fromNumber)
	msgData.Set("Body",msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	//Create the HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
	return nil
}
