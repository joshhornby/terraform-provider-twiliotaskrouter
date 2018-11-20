package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceTaskRouterActivity() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"available": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"true",
					"false",
				}, true),
				ForceNew: true,
			},
		},
		Create: activityCreate,
		Read:   activityRead,
		Update: activityUpdate,
		Delete: activityDelete,
	}
}
