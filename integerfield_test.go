package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testIntegerObject struct {
	Integer int `gforms:"integer"`
}

func TestIntegerField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewIntegerField("integer", nil),
	))
	data := url.Values{"integer": {"100"}}
	req, _ := http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		v, ok := form.CleanedData["integer"]
		if !ok {
			t.Error(`"integer" is required.`)
			return
		}
		_, ok = v.(int)
		if !ok {
			t.Error(`"integer" should be integer type.`)
			return
		}
		obj := new(testIntegerObject)
		form.MapTo(obj)
		if obj.Integer == 0 {
			t.Error(`"obj.Integer" should not be zero.`)
		}
	} else {
		t.Error("validation error.")
	}
}
