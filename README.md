# Sensu Go Twilio Handler Plugin
[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Twilio%20Handler-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-twilio-handler.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-twilio-handler)

This handler enables Sensu Go users to send SMS alerts using Twilio. There are some prerequisites that are required in order to be able to use this handler. 

1. You must have a Twilio account
2. You must obtain an [auth token](https://support.twilio.com/hc/en-us/articles/223136027-Auth-Tokens-and-How-to-Change-Them)
3. You must obtain an account SID
4. You must have a Twilio phone number

Once you have all of the above, you should be all set to use this handler. Read on to find more about how to use this handler.

## Installation

### As an asset

Assets are the best way to make use of this handler. If you're not using an asset, please consider doing so! You can find this asset on the[Bonsai Asset Index](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler).

You can download the asset definition there, or you can do a little bit of copy/pasta and use the one below:

```json
{
  "type": "Asset",
  "api_version": "core/v2",
  "metadata": {
    "name": "sensu-go-twilio-handler",
    "namespace": "CHANGEME",
    "labels": {
    },
    "annotations": {
    }
  },
  "spec": {
    "url": "https://github.com/asachs01/sensu-go-twilio-handler/releases/download/0.0.1/sensu-go-twilio-handler_0.0.1_linux_amd64.tar.gz",
    "sha512": "d054f14570069301dedf600e58ca7df78bd74e83377a44fee969a898e75c40ce1a30ee7eb24ce1a1c7f31c820a84e33b74cfb5b69163af22a45d6745eae780f0",
    "filters": [
      "entity.system.os == 'linux'",
      "entity.system.arch == 'amd64'"
    ]
  }
}
```

```yaml
---
type: Asset
api_version: core/v2
metadata:
  name: sensu-go-twilio-handler
  namespace: CHANGEME
  labels: {}
  annotations: {}
spec:
  url: https://github.com/asachs01/sensu-go-twilio-handler/releases/download/0.0.1/sensu-go-twilio-handler_0.0.1_linux_amd64.tar.gz
  sha512: d054f14570069301dedf600e58ca7df78bd74e83377a44fee969a898e75c40ce1a30ee7eb24ce1a1c7f31c820a84e33b74cfb5b69163af22a45d6745eae780f0
  filters:
  - entity.system.os == 'linux'
  - entity.system.arch == 'amd64'
```

**NOTE**: PLEASE ENSURE YOU UPDATE YOUR URL AND SHA512 BEFORE USING THE ASSET. If you don't, you might just be stuck on a super old version. Don't say I didn't warn you ¯\\_(ツ)_/¯

### Building from source

While it's generally recommended to use an asset, you can download a copy of the handler plugin from [releases][1],
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
  "type": "Handler",
  "api_version": "core/v2",
  "metadata": {
    "name": "sensu-go-twilio-handler",
    "namespace": "CHANGEME"
  },
  "spec": {
    "command": "sensu-go-twilio-handler -f +18558675309 -r +18559990210",
    "env_vars": [
      "TWILIO_ACCOUNT_SID=ACCTSIDXXXXXXXXXXXX",
      "TWILIO_AUTH_TOKEN=AUTHTOKENXXXXXXXXXX"
    ],
    "filters": [
      "is_incident",
      "not_silenced"
    ],
    "runtime_assets": [
      "sensu-go-twilio-handler"
    ],
    "timeout": 0,
    "type": "pipe"
  }
```
**YAML**
```yaml
type: Handler
api_version: core/v2
metadata:
  name: sensu-go-twilio-handler
  namespace: CHANGEME
spec:
  command: sensu-go-twilio-handler -f +18558675309 -r +18559990210
  env_vars:
  - TWILIO_ACCOUNT_SID=ACCTSIDXXXXXXXXXXXX
  - TWILIO_AUTH_TOKEN=AUTHTOKENXXXXXXXXXX
  filters:
  - is_incident
  - not_silenced
  runtime_assets:
  - sensu-go-twilio-handler
  timeout: 0
  type: pipe
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

You should then receive a message to your phone with the output of the error.

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/asachs01/sensu-go-twilio-handler/releases
