[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Twilio%20Handler-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-twilio-handler.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-twilio-handler)

# Sensu Go Twilio Handler Plugin

- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Asset definition](#asset-configuration)
  - [Handler definition](#resource-configuration)
- [Functionality](#functionality)
- [Installation from source and contributing](#installation-from-source-and-contributing)
- [Additional notes](#additional-notes)

## Overview

This handler enables Sensu Go users to send SMS alerts using Twilio. There are some prerequisites that are required in order to be able to use this handler. 

1. You must have a Twilio account
2. You must obtain an [auth token](https://support.twilio.com/hc/en-us/articles/223136027-Auth-Tokens-and-How-to-Change-Them)
3. You must obtain an account SID
4. You must have a Twilio phone number

Once you have all of the above, you should be all set to use this handler. Read on to find more about how to use this handler.


## Usage Examples

While this handler is meant to be run as part of a Sensu Go deployment, you can always run the binary directly to see what the options are, as well as test it to ensure that it work as expected. 

### Command line help

```text
The Sensu Go Handler for Twilio

Usage:
  sensu-go-twilio-handler [flags]

Flags:
  -s, --accountSid  string   The account SID for your Twilio account, uses the environment variable TWILIO_ACCOUNT_SID by default
  -t, --authToken   string   The authorization token for your Twilio account, uses the environment variable TWILIO_AUTH_TOKEN by default
  -f, --fromNumber  string   Your Twilio phone number
  -r, --recipients  string   The recipient phone number(s), separated by commas
  -h, --help                 help for sensu-go-twilio-handler
```

### Testing

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

![twilio alert screenshot](https://p299.p4.n0.cdn.getcloudapp.com/items/5zuJew4X/Screenshot_20190717-232020.png?v=f0244404f5516401e54d2ab64df1533d)

### Using the handler with one or more contacts

It is possible to use this plugin with multiple contacts, rather than a single number. This pattern is what is known in Sensu parlance as "contact routing." More will be added here as Sensu documentation is updated to more accurately describe this pattern.

## Configuration

### Asset registration

Assets are the best way to make use of this handler. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add asachs01/sensu-go-twilio-handler`

If you're using an earlier version of sensuctl, you can find this asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler).

### Asset definition

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

**NOTE**: ***PLEASE ENSURE YOU UPDATE YOUR URL AND SHA512 BEFORE USING THE ASSET!*** If you don't, you might just be stuck on a super old version. Don't say I didn't warn you ¯\\_(ツ)_/¯

### Handler definition

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

## Installation from Source and Contributing

While it's generally recommended to use an asset, you can download a copy of the handler plugin from [releases][1],
or create an executable script from this source.

From the local path of the sensu-go-twilio-handler repository:

```
go build -o /usr/local/bin/sensu-go-twilio-handler main.go
```

See [CONTRIBUTING][2] for more information on how to contribute to this plugin.

## Additional notes
### Supported operating systems

Support for this asset is as follows:

* Linux 64 bit
* Arm 64 bit

This is mostly due to the fact that this plugin is a handler. If you have the Sensu Go backend running on a different platform, [open an issue](https://github.com/asachs01/sensu-go-twilio-handler/issues/new) and provide some further details about the platform you'd like to see supported.


[1]: https://github.com/asachs01/sensu-go-twilio-handler/releases
[2]: https://github.com/asachs01/sensu-go-twilio-handler/blob/master/CONTRIBUTING.md

