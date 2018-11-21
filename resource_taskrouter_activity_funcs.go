package main

import (
	"net/url"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform/helper/schema"
)

func activityCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}
	data.Set("FriendlyName", d.Get("friendly_name").(string))
	data.Set("Available", d.Get("available").(string))

	activity, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Activities.Create(context.Background(), data)

	if err != nil {
		return err
	}

	d.SetId(activity.Sid)

	return activityUpdate(d, meta)
}

func activityUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}

	if d.HasChange("friendly_name") {
		data.Set("FriendlyName", d.Get("friendly_name").(string))
	}

	_, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Activities.Update(context.Background(), d.Id(), data)

	if err != nil {
		return err
	}

	return activityRead(d, meta)
}

func activityRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	activity, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Activities.Get(context.Background(), d.Id())
	if err != nil {
		return err
	}

	if activity == nil {
		d.SetId("")
		return err
	}

	d.SetId(activity.Sid)

	return nil
}

func activityDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Activities.Delete(context.Background(), d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
