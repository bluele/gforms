package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testBooleanObject struct {
	Check bool `gforms:"check"`
}

func TestTrueBooleanField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewBooleanField("check", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"check": {""}}.Encode()))
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
		obj := new(testBooleanObject)
		form.MapTo(obj)
		if obj.Check == false {
			t.Error(`"obj.Check" should not be false.`)
		}
	} else {
		t.Error("validation error.")
	}
}

func TestFalseBooleanField(t *testing.T) {
	Form := DefineForm(NewFields(
		NewBooleanField("check", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
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
		obj := new(testBooleanObject)
		form.MapTo(obj)
		if obj.Check == true {
			t.Error(`"obj.Check" should not be true.`)
		}
	} else {
		t.Error("validation error.")
	}
}

func TestBooleanFieldDefaultRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewBooleanField("check", nil),
	))
	form := Form()
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check">` {
		t.Errorf(`Incorrect HTML rendered for default boolean field: %s`, html)
		return
	}
}

func TestBooleanFieldInitialFalseRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewBooleanField("check", nil),
	))
	form := Form()
	field, _ := form.GetField("check")
	field.SetInitial("false")
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check">` {
		t.Errorf(`Incorrect HTML rendered for SetInitial("false") boolean field: %s`, html)
		return
	}
}

func TestBooleanFieldInitialTrueRender(t *testing.T) {
	Form := DefineForm(NewFields(
		NewBooleanField("check", nil),
	))
	form := Form()
	field, _ := form.GetField("check")
	field.SetInitial("true")
	html := strings.TrimSpace(form.Html())
	if html != `<input type="checkbox" name="check" checked>` {
		t.Errorf(`Incorrect HTML rendered for SetInitial("true") boolean field: %s`, html)
		return
	}
}