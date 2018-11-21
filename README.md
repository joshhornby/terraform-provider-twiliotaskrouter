# Terraform Provider

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.11.x

## Usage

- Download the latest release from [the releases page](https://github.com/joshhornby/terraform-provider-twiliotaskrouter/releases)
- `tar -xvzf terraform-provider-twiliotaskrouter_VERSION.tgz`
- `mkdir -p ~/.terraform.d/plugins/darwin_amd64 && cp terraform-provider-twiliotaskrouter ~/.terraform.d/plugins/darwin_amd64`

## Example

```
provider "twiliotaskrouter" {
  account_sid = "ACXXX"
  auth_token = "XXX"
  workspace_sid = "WSXXX"
}

resource "twiliotaskrouter_activity" "break_activity" {
    friendly_name = "Break"
    available = "false"
}

resource "twiliotaskrouter_activity" "online_activity" {
    friendly_name = "Online"
    available = "false"
}

resource "twiliotaskrouter_queue" "english_languages_queue" {
    friendly_name = "English Queue"
    reservation_activity_sid = "${twiliotaskrouter_activity.online_activity.id}"
    assignment_activity_sid = "${twiliotaskrouter_activity.break_activity.id}"
    target_workers = "languages HAS \"english\""

    depends_on = ["twiliotaskrouter_activity.break_activity", "twiliotaskrouter_activity.online_activity"]
}

resource "twiliotaskrouter_queue" "sales_queue" {
    friendly_name = "Sales Queue"
    reservation_activity_sid = "${twiliotaskrouter_activity.online_activity.id}"
    assignment_activity_sid = "${twiliotaskrouter_activity.break_activity.id}"
    target_workers = "skills HAS \"sales\""

    depends_on = ["twiliotaskrouter_activity.break_activity", "twiliotaskrouter_activity.online_activity"]
}

resource "twiliotaskrouter_workflow" "english_agents_workflow" {
    friendly_name = "Example Workflow"
    task_reservation_timeout = "20"
    configuration = <<EOF
{
"task_routing": {
        "default_filter": {
            "queue": "${twiliotaskrouter_queue.english_languages_queue.id}"
        },
        "filters": [
            {
                "filter_friendly_name": "Hello World",
                "expression": "type=='sales'",
                "targets": [
                    {
                        "queue": "${twiliotaskrouter_queue.english_languages_queue.id}",
                        "priority": "10",
                        "timeout": "15"
                    },
                    {
                        "queue": "${twiliotaskrouter_queue.sales_queue.id}",
                        "priority": "10"
                    }
                ]
            }
        ]
    }
}
EOF

  depends_on = ["twiliotaskrouter_queue.english_languages_queue", "twiliotaskrouter_queue.sales_queue"]
}

resource "twiliotaskrouter_worker" "my_worker" {
    friendly_name = "John Doe"
    attributes = <<EOF
{
"type": "sales"
}
EOF
}

```

The provider configuration block accepts the following arguments:

- `account_sid` - (Required) Your SID (application ID) for the the Twilio API. May alternatively be set via the
  `TWILIO_SID` environment variable.
- `auth_token` - (Required) The API auth token to use when making requests. May alternatively
  be set via the `TWILIO_AUTH_TOKEN` environment variable.
- `workspace_sid` - (Required) The sid of your Task Router Workspace. May alternatively
  be set via the `TWILIO_WORKSPACE_SID` environment variable.

### With Thanks

- Thanks To Kevin Burke for his help over on the [Twilio Go library](https://github.com/kevinburke/twilio-go)
- Inspired by https://github.com/tulip/terraform-provider-twilio
