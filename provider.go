package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	twilioc "github.com/kevinburke/twilio-go"
)

func provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_SID", nil),
				Description: "The SID (application ID) for the the Twilio API",
			},
			"auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "Oauth token for the the Twilio API",
			},
			"workspace_sid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_WORKSPACE_SID", nil),
				Description: "Workspace SID for Task Router",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_taskrouter_activity": resourceTaskRouterActivity(),
			"twilio_taskrouter_queue":    resourceTaskRouterQueue(),
			"twilio_taskrouter_workflow": resourceTaskRouterWorkflow(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type twilioMeta struct {
	Client       *twilioc.Client
	WorkspaceSid string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &twilioMeta{
		Client: twilioc.NewClient(
			d.Get("account_sid").(string),
			d.Get("auth_token").(string),
			nil,
		),
		WorkspaceSid: d.Get("workspace_sid").(string),
	}, nil
}
