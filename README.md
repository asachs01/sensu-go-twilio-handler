# Sensu Go Twilio Handler Plugin
[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Twilio%20Handler-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-twilio-handler) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-twilio-handler.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-twilio-handler)

This handler enables Sensu Go users to send SMS alerts using Twilio.

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
        "namespace": "default",
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
  namespace: default
  name: sensu-go-twilio-handler
spec:
  "...": "..."
```

## Usage Examples

Help:

```
The Sensu Go Handler for Twilio

Usage:
  sensu-go-twilio-handler [flags]

Flags:
  -f, --foo string   example
  -h, --help         help for sensu-go-twilio-handler
```

## Contributing

See https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/asachs01/sensu-go-twilio-handler/releases
