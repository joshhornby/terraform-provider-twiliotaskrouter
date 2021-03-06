package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceTaskRouterWorkflow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration": &schema.Schema{
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: suppressEquivalentJsonDiffs,
			},
			"assignment_callback_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_assignment_callback_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_reservation_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: taskRouterWorkflowCreate,
		Read:   taskRouterWorkflowRead,
		Update: taskRouterWorkflowUpdate,
		Delete: taskRouterWorkflowDelete,
	}
}
