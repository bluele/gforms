package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testNullBooleanObject struct {
	Check bool `gforms:"check"`
}

func TestTrueNullBooleanField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"check": {"true"}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		v, ok := form.CleanedData["check"]
		if !ok {
			t.Error(`"check" is required.`)
			return
		}
		_, ok = v.(bool)
		if !ok {
			t.Error(`"check" should be boolean type.`)
			return
		}
		obj := new(testNullBooleanObject)
		form.MapTo(obj)
		if obj.Check == false {
			t.Error(`"obj.Check" should not be false.`)
		}
	} else {
		t.Error("validation error.")
	}
}

func TestFalseNullBooleanField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"check": {"false"}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		v, ok := form.CleanedData["check"]
		if !ok {
			t.Error(`"check" is required.`)
			return
		}
		_, ok = v.(bool)
		if !ok {
			t.Error(`"check" should be boolean type.`)
			return
		}
		obj := new(testNullBooleanObject)
		form.MapTo(obj)
		if obj.Check == false {
			t.Error(`"obj.Check" should be false.`)
		}
	} else {
		t.Error("validation error.")
	}
}

func TestFalseNullBooleanFieldEmpty(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if form.IsValid() {
		_, ok := form.CleanedData["check"]
		if ok {
			t.Error(`"check" should not exist.`)
			return
		}
	}
}

func TestTrueNullBooleanFieldJsonRequired(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", Validators{Required()}),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader("{}"))
	req.Header.Add("Content-Type", "application/json")
	form := Form(req)
	if form.IsValid() {
		t.Error("Null boolean field should be required.")
	}
}

func TestNullBooleanFieldDefaultRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	form := Form()
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check">` {
		t.Errorf(`Incorrect HTML rendered for default null boolean field: %s`, html)
		return
	}
}

func TestNullBooleanFieldInitialFalseRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	form := Form()
	field, _ := form.GetField("check")
	field.SetInitial("false")
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check">` {
		t.Errorf(`Incorrect HTML rendered for SetInitial("false") null boolean field: %s`, html)
		return
	}
}

func TestNullBooleanFieldInitialTrueRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewNullBooleanField("check", nil),
	))
	form := Form()
	field, _ := form.GetField("check")
	field.SetInitial("true")
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check" checked>` {
		t.Errorf(`Incorrect HTML rendered for SetInitial("true") null boolean field: %s`, html)
		return
	}
}