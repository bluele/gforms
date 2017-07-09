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

func testEmail(t *testing.T, email string, expectValid bool) {
	Form := DefineForm(NewFields(
		NewTextField(
			"email",
			Validators{
				EmailValidator(),
			},
		),
	))
	
	req, _ := http.NewRequest("POST", "/", strings.NewReader(url.Values{"email": {email}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	form := Form(req)
	if expectValid {
		if !form.IsValid() {
			t.Error("Not expected: validation error for email " + email)
		}	
	} else {
		if form.IsValid() {
			t.Error("Expected: validation error for email " + email)
		}
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
	
	// Normal email
	testEmail(t, "junkxdev@gmail.com", true)
	
	// Normal email with special characters
	testEmail(t, "user.1234!#$%&'*+-/=?^_`{|}~ ()@gmail.com", true)
	
	// Email missing '.' after '@' - rare and not encouraged but still legal.
	testEmail(t, "junkxdev@gmail", true)
	
	// Bad email - missing '@'
	testEmail(t, "junkxdevgmail.com", false)
	testEmail(t, "@gmail.com", false)
	
	// Valid UTF8 emails
	testEmail(t, "闪闪发光@闪闪发光.com", true)
	testEmail(t, "Pelé@example.com", true) // Latin alphabet (with diacritics)
	testEmail(t, "δοκιμή@παράδειγμα.δοκιμή", true) // Greek alphabet
	testEmail(t, "我買@屋企.香港", true) // Traditional Chinese characters
	testEmail(t, "甲斐@黒川.日本", true) // Japanese characters
	testEmail(t, "чебурашка@ящик-с-апельсинами.рф", true) // Cyrillic characters
	testEmail(t, "संपर्क@डाटामेल.भारत", true) // Hindi email address

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
