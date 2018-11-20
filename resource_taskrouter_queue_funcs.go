package main

import (
	"net/url"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform/helper/schema"
)

func taskRouterQueueCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}
	data.Set("FriendlyName", d.Get("name").(string))
	data.Set("ReservationActivitySid", d.Get("reservation_activity_sid").(string))
	data.Set("AssignmentActivitySid", d.Get("assignment_activity_sid").(string))
	data.Set("TargetWorkers", d.Get("target_workers").(string))

	taskQueue, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Queues.Create(context.Background(), data)

	if err != nil {
		return err
	}

	d.SetId(taskQueue.Sid)

	return taskRouterQueueUpdate(d, meta)
}

func taskRouterQueueUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}

	if d.HasChange("name") {
		data.Set("FriendlyName", d.Get("name").(string))
	}

	if d.HasChange("reservation_activity_sid") {
		data.Set("ReservationActivitySid", d.Get("reservation_activity_sid").(string))
	}

	if d.HasChange("assignment_activity_sid") {
		data.Set("AssignmentActivitySid", d.Get("assignment_activity_sid").(string))
	}

	if d.HasChange("target_workers") {
		data.Set("TargetWorkers", d.Get("target_workers").(string))
	}

	_, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Queues.Update(context.Background(), d.Id(), data)

	if err != nil {
		return err
	}

	return taskRouterQueueRead(d, meta)
}

func taskRouterQueueRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	taskQueue, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Queues.Get(context.Background(), d.Id())
	if err != nil {
		return err
	}

	if taskQueue == nil {
		d.SetId("")
		return err
	}

	d.SetId(taskQueue.Sid)

	return nil
}

func taskRouterQueueDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Queues.Delete(context.Background(), d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
