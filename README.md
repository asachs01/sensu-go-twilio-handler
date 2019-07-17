# Sensu Go Twilio Handler Plugin
[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Twilio%20Handler-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-twilio-handler.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-twilio-handler)

This handler enables Sensu Go users to send SMS alerts using Twilio. There are some prerequisites that are required in order to be able to use this handler. 

1. You must have a Twilio account
2. You must obtain an [auth token](https://support.twilio.com/hc/en-us/articles/223136027-Auth-Tokens-and-How-to-Change-Them)
3. You must obtain an account SID
4. You must have a Twilio phone number

Pro

TODO: SCREENSHOTS

## Installation

Download the latest version of the sensu-go-twilio-handler from [releases][1],
or create an executable script from this source.

From the local path of the sensu-go-twilio-handler repository:

```
go build -o /usr/local/bin/sensu-go-twilio-handler main.go
```

## Configuration

Example Sensu Go definition:
**JSON**
```json
{
    "api_version": "core/v2",
    "type": "Handler",
    "metadata": {
        "namespace": "CHANGEME",
        "name": "sensu-go-twilio-handler"
    },
    "spec": {
        "...": "..."
    }
}
```
**YAML**
```yaml
---
api_version: core/v2
type: Handler
metadata:
  namespace: CHANGEME
  name: sensu-go-twilio-handler
spec:
  "...": "..."
```

## Usage Examples

Help:

```text
The Sensu Go Handler for Twilio

Usage:
  sensu-go-twilio-handler [flags]

Flags:
  -s, --accountSid  string   The account SID for your Twilio account, uses the environment variable TWILIO_ACCOUNT_SID by default
  -t, --authToken   string   The authorization token for your Twilio account, uses the environment variable TWILIO_AUTH_TOKEN by default
  -f, --fromNumber  string   Your Twilio phone number
  -r, --recipient   string   The recipient's phone number
  -h, --help         help for sensu-go-twilio-handler
```

## Testing

To test and see if this handler works, do the following:

Clone the repo:
```
git clone github.com/asachs01/sensu-go-twilio-handler
```

Run the following command:
```
cat example-event.json | ./sensu-go-twilio-handler  -s ACCTSIDXXXXXXXXXXXX -t AUTHTOKENXXXXXXXXXX  -f +18558675309 -r +18559990210
```

You should receive a message to your phone with the output of the error

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/asachs01/sensu-go-twilio-handler/releases
