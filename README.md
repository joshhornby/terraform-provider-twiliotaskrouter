# Terraform Provider

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/joshhornby/terraform-provider-twiliotaskrouter`

```sh
$ mkdir -p $GOPATH/src/github.com/joshhornby; cd $GOPATH/src/github.com/joshhornby
$ git clone git@github.com:joshhornby/terraform-provider-twiliotaskrouter.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/joshhornby/terraform-provider-twiliotaskrouter
$ make build
```

The provider configuration block accepts the following arguments:

- `account_sid` - (Required) Your SID (application ID) for the the Twilio API. May alternatively be set via the
  `TWILIO_SID` environment variable.

- `auth_token` - (Required) The API auth token to use when making requests. May alternatively
  be set via the `TWILIO_AUTH_TOKEN` environment variable.
