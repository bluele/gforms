# GForms
A flexible forms validation and rendering library for golang web development. 
Inspired by [django-forms](https://docs.djangoproject.com/en/dev/topics/forms/) and [wtforms](https://github.com/wtforms/wtforms).

[![wercker status](https://app.wercker.com/status/51a7f6720baf8e67a28241790380d19b/s "wercker status")](https://app.wercker.com/project/bykey/51a7f6720baf8e67a28241790380d19b)

## Overview

* Validate HTTP request
* Rendering form-html
* Support parsing content-type "form-urlencoded", "json"
* Support many widgets for form field.

# Getting Started

## Install

```
go get github.com/bluele/gforms
```

## Examples

See [examples](https://github.com/bluele/gforms/tree/master/examples).

## Usage

### Define Form

```go
userForm := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    gforms.Validators{
      gforms.Required(),
      gforms.MaxLengthValidator(32),
    },
  ),
  gforms.NewFloatField(
    "weight",
    gforms.Validators{},
  ),
))
```

### Validate HTTP request

Server ([code](https://github.com/bluele/gforms/blob/master/examples/simple_form.go)):

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func main() {
  tplText := `<form method="post">
{{range $i, $field := .Fields}}
  <label>{{$field.GetName}}: </label>{{$field.Html}}
  {{range $ei, $err := $field.Errors}}<label class="error">{{$err}}</label>{{end}}<br />
{{end}}<input type="submit">
</form>
  `
  tpl := template.Must(template.New("tpl").Parse(tplText))

  userForm := gforms.DefineForm(gforms.NewFields(
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(),
        gforms.MaxLengthValidator(32),
      },
    ),
    gforms.NewFloatField(
      "weight",
      gforms.Validators{},
    ),
  ))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    form := userForm(r)
    if r.Method != "POST" {
      tpl.Execute(w, form)
      return
    }
    if !form.IsValid() {
      tpl.Execute(w, form)
      return
    }
    user := User{}
    form.MapTo(&user)
    fmt.Fprintf(w, "ok: %v", user)
  })
  http.ListenAndServe(":9000", nil)
}
```

Client:

```
# show html form
$ curl -X GET localhost:9000/users
<form method="post">
  <label>name: </label><input type="text" name="name" value=""></input>
  <br />

  <label>weight: </label><input type="text" name="weight" value=""></input>
  <br />
<input type="submit">
</form>

# valid request
$ curl -X POST localhost:9000/users -d 'name=bluele&weight=71.9'
ok: {bluele 71.9}

# "name" field is required.
$ curl -X POST localhost:9000/users -d 'weight=71.9'
<form method="post">
  <label>name: </label><input type="text" name="name" value=""></input>
  <label class="error">This field is required.</label><br />

  <label>weight: </label><input type="text" name="weight" value="71.9"></input>
  <br />
<input type="submit">
</form>
```

### Define Form by Struct Model

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func initForm() {
  userForm := gforms.DefineModelForm(User{}, gforms.NewFields(
    // override User.name field
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(),
        gforms.MaxLengthValidator(32),
      },
    ),
  ))
  /* equal an above defined form.
  userForm := gforms.DefineForm(gforms.NewFields(
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(),
        gforms.MaxLengthValidator(32),
      },
    ),
    gforms.NewFloatField(
      "weight",
      gforms.Validators{},
    ),
  ))
  */
}
```

## Render HTML

### FormInstance#Html

```go
form := userForm(r)
fmt.Println(form.Html())
/* 
# Output
<input type="text" name="name"></input>
<input type="text" name="weight"></input>
*/
```

### FieldInstance#Html

```
form := userForm(r)
fmt.Println(form.GetField("name").Html())
/* 
# Output
<input type="text" name="name"></input>
*/
```

## Parse request data

### (Default) Parse `*http.Request` to create a new `FormInstance`.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  form := userForm(r)  
  ...
}
```

### Parse `net/url.Values` to create a new `FormInstance`.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // parse querystring values.
  form := userForm.FromUrlValues(r.URL.Query())
  ...
}
```

## Customize Formfield attributes

```go
customForm := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    gforms.Validators{
      gforms.Required(),
    },
    gforms.TextInputWidget(
      map[string]string{
        "class": "custom",
      },
    )),
))

form := customForm(r)
fmt.Println(form.Html())
/* 
# Output
<input type="text" name="name" value="" class="custom"></input>
*/
```

## Custom Validation error message

```go
userForm := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    gforms.Validators{
      gforms.Required("Custom error required message."),
      gforms.MaxLengthValidator(32, "Custom error maxlength message."),
    },
  ),
))
```

## Support Fields

### TextField

It maps value to FormInstance.CleanedData as type `string`.

```go
gforms.NewTextField(
  "text",
  gforms.Validators{},
)
```

### BooleanField

It maps value to FormInstance.CleanedData as type `bool`.

```go
gforms.NewBooleanField(
  "checked",
  gforms.Validators{},
)
```

### IntegerField

It maps value to FormInstance.CleanedData as type `int`.

```go
gforms.NewIntegerField(
  "number",
  gforms.Validators{},
)
```

### FloatField

It maps value to FormInstance.CleanedData as type `float32` or `float64`.

```go
gforms.NewFloatField(
  "floatNumber",
  gforms.Validators{},
)
```

### MultipleTextField

It maps value to FormInstance.CleanedData as type `[]string`.

```go
gforms.NewMultipleTextField(
  "texts",
  gforms.Validators{},
)
```

### DateTimeField

It maps value to FormInstance.CleanedData as type `time.Time`.

```go
gforms.NewDateTimeField(
  "date", 
  DefaultDateTimeFormat, 
  gforms.Validators{},  
)
```

## Support Validators

### Required validator

Added an error msg to FormInstance.Errors() if the field is not provided.

```go
gforms.Validators{
  gforms.Required(),
},
```

### Regexp validator

Added an error msg to FormInstance.Errors() if the regexp doesn't matched a value.

```go
gforms.Validators{
  gforms.RegexpValidator(`number-\d+`),
},
```

### Email validator

Added an error msg to FormInstance.Errors() if a value doesn't looks like an email address.

```go
gforms.Validators{
  gforms.EmailValidator(),
},
```

### URL Validator

Added an error msg to FormInstance.Errors() if a value doesn't looks like an url.

```go
gforms.Validators{
  gforms.URLValidator(),
},
```

### MinLength Validator

Added an error msg to FormInstance.Errors() if the length of value is less than specified first argument.

```go
gforms.Validators{
  gforms.MinLengthValidator(16),
},
```

### MaxLength Validator

Added an error msg to FormInstance.Errors() if the length of value is greater than specified first argument.

```go
gforms.Validators{
  gforms.MaxLengthValidator(256),
},
```

### MinValueValidator

Added an error msg to FormInstance.Errors() if value is less than specified first argument.

```go
gforms.Validators{
  gforms.MinValueValidator(16),
},
```

### MaxValueValidator

Added an error msg to FormInstance.Errors() if value is greater than specified first argument.

```go
gforms.Validators{
  gforms.MaxValueValidator(256),
},
```

## Support Widgets

### SelectWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "gender",
    gforms.Validators{
      gforms.Required(),
    },
    gforms.SelectWidget(
      map[string]string{
        "class": "custom",
      },
      func() gforms.SelectOptions {
        return gforms.StringSelectOptions([][]string{
          {"Men", "0"},
          {"Women", "1"},
        })
      },
    ),
  ),
))

form = Form()
fmt.Println(form.Html())
/*
# output
<select class="custom">
<option value="0">Men</option>
<option value="1">Women</option>
</select>
*/
```

### RadioSelectWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
    gforms.NewTextField(
      "lang",
      gforms.Validators{
        gforms.Required(),
      },
      gforms.RadioSelectWidget(
        map[string]string{
          "class": "custom",
        },
        func() gforms.RadioOptions {
          return gforms.StringRadioOptions([][]string{
            {"Golang", "0", "false", "false"},
            {"Python", "1", "false", "true"},
          })
        },
      ),
    ),  
))

form = Form()
fmt.Println(form.Html())
/*
# output
<input type="radio" name="lang" value="0">Golang
<input type="radio" name="lang" value="1" disabled>Python
*/
```

### CheckboxMultipleWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
    gforms.NewMultipleTextField(
      "lang",
      gforms.Validators{
        gforms.Required(),
      },
      gforms.CheckboxMultipleWidget(
        map[string]string{
          "class": "custom",
        },
        func() gforms.CheckboxOptions {
          return gforms.StringCheckboxOptions([][]string{
            {"Golang", "0", "false", "false"},
            {"Python", "1", "false", "true"},
          })
        },
      ),
    ),
))

form := Form()
fmt.Println(form.Html())
/*
# output
<input type="checkbox" name="lang" value="0">Golang
<input type="checkbox" name="lang" value="1" disabled>Python
*/
```

# TODO

* Support FileField, DateField, DateTimeField
* Writing more godoc and unit tests.
* Improve performance.

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
