package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

type testDateTimeObject struct {
	Date time.Time `gforms:"date"`
}

func TestInputValidDateFormat(t *testing.T) {
	Form := DefineForm(NewFields(
		NewDateTimeField("date", DefaultDateFormat, nil),
	))
	reqDate := "2014-06-18"
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"date": {reqDate}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	form := Form(req)
	if !form.IsValid() {
		t.Error("validation error.")
		return
	}
	v, ok := form.CleanedData["date"]
	if !ok {
		t.Error(`"date" is required.`)
	}

	_, ok = v.(time.Time)
	if !ok {
		t.Error(`"date" should be time.Time type.`)
		return
	}
	obj := new(testDateTimeObject)
	form.MapTo(obj)
	if obj.Date.Year() != 2014 || obj.Date.Month() != 6 || obj.Date.Day() != 18 {
		t.Error("exptected: " + reqDate)
	}
}

func TestInputInValidDateFormat(t *testing.T) {
	Form := DefineForm(NewFields(
		NewDateTimeField("date", DefaultDateFormat, nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"date": {"20140618"}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	form := Form(req)
	if form.IsValid() {
		t.Error("expected: validation error.")
		return
	}
}

func TestCustomDateFormat(t *testing.T) {
	Form := DefineForm(NewFields(
		NewDateTimeField("date", "2006/01/02", nil),
	))
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"date": {"2014/06/18"}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	form := Form(req)
	if !form.IsValid() {
		t.Error("not expected: validation error.")
		return
	}
}

func TestInputValidDateTimeFormat(t *testing.T) {
	Form := DefineForm(NewFields(
		NewDateTimeField("date", DefaultDateTimeFormat, nil),
	))
	reqDate := "2014-06-18 05:14:16"
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"date": {reqDate}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	form := Form(req)
	if !form.IsValid() {
		t.Error("validation error.")
		return
	}
	v, ok := form.CleanedData["date"]
	if !ok {
		t.Error(`"date" is required.`)
	}

	_, ok = v.(time.Time)
	if !ok {
		t.Error(`"date" should be time.Time type.`)
		return
	}
	obj := new(testDateTimeObject)
	form.MapTo(obj)
	if obj.Date.Year() != 2014 || obj.Date.Month() != 6 || obj.Date.Day() != 18 || obj.Date.Hour() != 5 || obj.Date.Minute() != 14 || obj.Date.Second() != 16 {
		t.Error("given: " + obj.Date.String() + ", exptected: " + DefaultDateTimeFormat)
	}
}
