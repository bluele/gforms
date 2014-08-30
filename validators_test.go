package gforms

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestRequiredValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"name",
			Validators{
				Required(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {"bluele"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}
}

func TestMaxLengthValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"name",
			Validators{
				MaxLengthValidator(8),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {"bluele"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {"abcdefghi"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestMinLengthValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"name",
			Validators{
				MinLengthValidator(4),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {"bluele"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"name": {"abc"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestRegexpValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"id",
			Validators{
				RegexpValidator(`^\d+$`),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"id": {"123"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"id": {"abc"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestEmailValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"email",
			Validators{
				EmailValidator(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"email": {"junkxdev@gmail.com"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"email": {"abc"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestURLValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewTextField(
			"url",
			Validators{
				URLValidator(),
			},
		),
	))

	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"url": {"https://github.com"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"url": {"abc"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestMaxValueValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewIntegerField(
			"value",
			Validators{
				MaxValueValidator(100),
			},
		),
	))
	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"value": {"100"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"value": {"101"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}

func TestMinValueValidator(t *testing.T) {
	Form := DefineForm(NewFields(
		NewIntegerField(
			"value",
			Validators{
				MinValueValidator(100),
			},
		),
	))
	req1, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"value": {"100"}}.Encode()))
	req1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form1 := Form(req1)
	if !form1.IsValid() {
		t.Error("Not expected: validation error.")
	}

	req2, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"value": {"99"}}.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form2 := Form(req2)
	if form2.IsValid() {
		t.Error("Expected: validation error.")
	}

	req3, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{}.Encode()))
	req3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form3 := Form(req3)
	if !form3.IsValid() {
		t.Error("Not expected: validation error.")
	}
}
