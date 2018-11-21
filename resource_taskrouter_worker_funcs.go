package main

import (
	"net/url"

	"golang.org/x/net/context"

	"github.com/hashicorp/terraform/helper/schema"
)

func workerCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}
	data.Set("FriendlyName", d.Get("friendly_name").(string))
	data.Set("Attributes", d.Get("attributes").(string))

	worker, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workers.Create(context.Background(), data)

	if err != nil {
		return err
	}

	d.SetId(worker.Sid)

	return workerUpdate(d, meta)
}

func workerUpdate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	data := url.Values{}

	if d.HasChange("friendly_name") {
		data.Set("FriendlyName", d.Get("friendly_name").(string))
	}

	if d.HasChange("attributes") {
		data.Set("Attributes", d.Get("attributes").(string))
	}

	_, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workers.Update(context.Background(), d.Id(), data)

	if err != nil {
		return err
	}

	return workerRead(d, meta)
}

func workerRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	worker, err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workers.Get(context.Background(), d.Id())
	if err != nil {
		return err
	}

	if worker == nil {
		d.SetId("")
		return err
	}

	d.SetId(worker.Sid)

	return nil
}

func workerDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*twilioMeta)

	err := m.Client.TaskRouter.Workspace(m.WorkspaceSid).Workers.Delete(context.Background(), d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
