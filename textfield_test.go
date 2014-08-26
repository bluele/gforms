package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testTextObject struct {
	Name string `gforms:"name"`
}

func TestTextField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField("name", nil),
	))

	data := url.Values{"name": {"bluele"}}
	req, _ := http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		v, ok := form.CleanedData["name"]
		if !ok {
			t.Error(`"name" is required.`)
			return
		}
		_, ok = v.(string)
		if !ok {
			t.Error(`"name" should be string type.`)
			return
		}
		obj := new(testTextObject)
		form.MapTo(obj)
		if obj.Name == "" {
			t.Error(`"obj.Name" should not be empty string.`)
		}
	} else {
		t.Error("validation error.")
	}
}
