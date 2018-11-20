package main

import (
	"net/url"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform/helper/schema"
)

func taskRouterWorkflowCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}
	data.Set("FriendlyName", d.Get("name").(string))
	data.Set("Configuration", d.Get("configuration").(string))
	data.Set("AssignmentCallbackUrl", d.Get("assignment_callback_url").(string))
	data.Set("FallbackAssignmentCallbackUrl", d.Get("fallback_assignment_callback_url").(string))
	data.Set("TaskReservationTimeout", d.Get("task_reservation_timeout").(string))

	workflow, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workflows.Create(context.Background(), data)

	if err != nil {
		return err
	}

	d.SetId(workflow.Sid)

	return taskRouterWorkflowUpdate(d, meta)
}

func taskRouterWorkflowUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}

	if d.HasChange("name") {
		data.Set("FriendlyName", d.Get("name").(string))
	}

	if d.HasChange("configuration") {
		data.Set("Configuration", d.Get("configuration").(string))
	}

	if d.HasChange("assignment_callback_url") {
		data.Set("AssignmentCallbackUrl", d.Get("assignment_callback_url").(string))
	}

	if d.HasChange("fallback_assignment_callback_url") {
		data.Set("FallbackAssignmentCallbackUrl", d.Get("fallback_assignment_callback_url").(string))
	}

	if d.HasChange("task_reservation_timeout") {
		data.Set("TaskReservationTimeout", d.Get("task_reservation_timeout").(string))
	}

	_, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workflows.Update(context.Background(), d.Id(), data)

	if err != nil {
		return err
	}

	return taskRouterWorkflowRead(d, meta)
}

func taskRouterWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	workflow, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workflows.Get(context.Background(), d.Id())
	if err != nil {
		return err
	}

	if workflow == nil {
		d.SetId("")
		return err
	}

	d.SetId(workflow.Sid)

	return nil
}

func taskRouterWorkflowDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workflows.Delete(context.Background(), d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
