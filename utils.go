package main

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
)

func suppressEquivalentJsonDiffs(k, old, new string, d *schema.ResourceData) bool {
	ob := bytes.NewBufferString("")
	if err := json.Compact(ob, []byte(old)); err != nil {
		return false
	}

	nb := bytes.NewBufferString("")
	if err := json.Compact(nb, []byte(new)); err != nil {
		return false
	}

	return jsonBytesEqual(ob.Bytes(), nb.Bytes())
}

func jsonBytesEqual(b1, b2 []byte) bool {
	var o1 interface{}
	if err := json.Unmarshal(b1, &o1); err != nil {
		return false
	}
	var o2 interface{}
	if err := json.Unmarshal(b2, &o2); err != nil {
		return false
	}
	return reflect.DeepEqual(o1, o2)
}
