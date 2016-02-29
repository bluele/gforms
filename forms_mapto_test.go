package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"time"
	"testing"
)

func TestSimpleMapToSuccess(t *testing.T) {
	testName := "bluele"
	Form := DefineForm(NewFields(
		NewTextField(
			"Name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"Name": {testName}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Name string
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Name != testName {
		t.Errorf("Expected Name == %s, got \"%s\".", testName, domain.Name)
	}
}

func TestSimpleMapToFieldNameMiss(t *testing.T) {
	testName := "bluele"
	Form := DefineForm(NewFields(
		NewTextField(
			"name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {testName}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Name string
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Name != "" {
		t.Errorf("Expected Name == \"\", got \"%s\".", domain.Name)
	}
}

func TestMapToPtr(t *testing.T) {
	testName := "bluele"
	Form := DefineForm(NewFields(
		NewTextField(
			"Name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"Name": {testName}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Name *string
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Name == nil {
		t.Error("Name should not be nil.")
		return
	}
	if *domain.Name != testName {
		t.Errorf("Expected Name == %s, got \"%s\".", testName, domain.Name)
	}
}

func TestMapToPtrFieldNameMiss(t *testing.T) {
	testName := "bluele"
	Form := DefineForm(NewFields(
		NewTextField(
			"name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {testName}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Name *string
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Name != nil {
		t.Errorf("Name should be nil, got \"%s\".", domain.Name)
		return
	}
}

func TestMapToDoublePtr(t *testing.T) {
	testName := "bluele"
	Form := DefineForm(NewFields(
		NewTextField(
			"Name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"Name": {testName}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Name **string
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Name == nil {
		t.Error("Name should not be nil.")
		return
	}
	if **domain.Name != testName {
		t.Errorf("Expected Name == %s, got \"%s\".", testName, domain.Name)
	}
}

func TestMapToStruct(t *testing.T) {
	testDate := "2016-02-19 11:22:33"
	dateFormat := "2006-01-02 15:04:05"
	testDateTime := time.Time{}
	testDateTime, err := time.Parse(dateFormat, testDate)
	if err != nil {
		t.Error("Not expected: date parsing error.")
	}

	Form := DefineForm(NewFields(
		NewDateTimeField(
			"Expiration",
			dateFormat,
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"Expiration": {testDate}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Expiration time.Time
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if !domain.Expiration.Equal(testDateTime) {
		t.Errorf("Expected Expiration == %s, got \"%s\".", testDateTime.String(), domain.Expiration.String())
	}
}

func TestMapToPtrToStruct(t *testing.T) {
	testDate := "2016-02-19 11:22:33"
	dateFormat := "2006-01-02 15:04:05"
	testDateTime := time.Time{}
	testDateTime, err := time.Parse(dateFormat, testDate)
	if err != nil {
		t.Error("Not expected: date parsing error.")
	}

	Form := DefineForm(NewFields(
		NewDateTimeField(
			"Expiration",
			dateFormat,
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"Expiration": {testDate}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	type Domain struct {
		Expiration *time.Time
	}

	domain := Domain{}
	form1.MapTo(&domain)
	if domain.Expiration == nil {
		t.Error("Expiration should not be nil.")
		return
	}
	if !(*domain.Expiration).Equal(testDateTime) {
		t.Errorf("Expected Expiration == %s, got \"%s\".", testDateTime.String(), domain.Expiration.String())
	}
}
