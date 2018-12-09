package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceTaskRouterWorker() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attributes": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: suppressEquivalentJsonDiffs,
			},
		},
		Create: workerCreate,
		Read:   workerRead,
		Update: workerUpdate,
		Delete: workerDelete,
	}
}
