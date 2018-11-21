package main

import (
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceTaskRouterQueue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reservation_activity_sid": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(WA)[a-zA-Z0-9]{32}$`), "Invalid reservation_activity_sid"),
			},
			"assignment_activity_sid": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(WA)[a-zA-Z0-9]{32}$`), "Invalid assignment_activity_sid"),
			},
			"target_workers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: taskRouterQueueCreate,
		Read:   taskRouterQueueRead,
		Update: taskRouterQueueUpdate,
		Delete: taskRouterQueueDelete,
	}
}
