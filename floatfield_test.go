package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testFloatObject struct {
	Float float64 `gforms:"float"`
}

func TestFloatField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewFloatField("float", nil),
	))
	data := url.Values{"float": {"100"}}
	req, _ := http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		v, ok := form.CleanedData["float"]
		if !ok {
			t.Error(`"float" is required.`)
			return
		}
		_, ok = v.(float64)
		if !ok {
			t.Error(`"float" should be float type.`)
			return
		}
		obj := new(testFloatObject)
		form.MapTo(obj)
		if obj.Float == 0.0 {
			t.Error(`"obj.Float" should not be zero.`)
		}
	} else {
		t.Error("validation error.")
	}
}
