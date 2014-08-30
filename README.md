# GForms
A flexible forms validation and rendering library for golang web development. (under heavy development)

[![wercker status](https://app.wercker.com/status/51a7f6720baf8e67a28241790380d19b/s "wercker status")](https://app.wercker.com/project/bykey/51a7f6720baf8e67a28241790380d19b)

## Overview

* Validate HTTP request
* Rendering form-html
* Support parsing content-type "form-urlencoded", "json"

# Getting Started

## Install

```
go get github.com/bluele/gforms
```

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

Server:

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func main() {
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

  http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    form := userForm(r)
    if r.Method != "POST" {
      fmt.Fprintf(w, form.Html())
      return
    }
    if form.IsValid() { // Validate request body
      user := User{}
      form.MapTo(&user)
      fmt.Fprintf(w, "%v", user)
    } else {
      fmt.Fprintf(w, "%v", form.Errors)
    }
  })
  http.ListenAndServe(":9000", nil)
}
```

Client:

```
$ curl -X GET localhost:9000/users
<input type="text" name="name"></input>
<input type="text" name="weight"></input>

$ curl -X POST localhost:9000/users -d 'name=bluele&weight=71.9'
{bluele 71.9}

# "name" field is required.
$ curl -X POST localhost:9000/users -d 'weight=71.9'
map[name:This field is required]

# also support json request
$ curl -X POST -H "Content-type: application/json" localhost:9000/users -d '{"name":"bluele", "weight":71.9}'
{bluele 71.9}
```

### Define Form by Struct Model

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func initForm() {
  userForm := gforms.DefineModelForm(gforms.NewFields(
    User{},
    // override User.name field
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(),
        gforms.MaxLengthValidator(32),
      },
    ),
  ))
}
```

### Validate HTTP request

Server:

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func main() {
  userForm := gforms.DefineModelForm(gforms.NewFields(
    User{},
    // override User.name field
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(),
        gforms.MaxLengthValidator(32),
      },
    ),
  ))

  http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    form := userForm(r)
    if r.Method != "POST" {
      fmt.Fprintf(w, form.Html())
      return
    }
    if form.IsValid() { // Validate request body
      user := form.GetModel().(User)
      fmt.Fprintf(w, "%v", user)
    } else {
      fmt.Fprintf(w, "%v", form.Errors)
    }
  })
  http.ListenAndServe(":9000", nil)
}
```

Client:

```
$ curl -X GET localhost:9000/users
<input type="text" name="name"></input>
<input type="text" name="weight"></input>

$ curl -X POST localhost:9000/users -d 'name=bluele&weight=71.9'
{bluele 71.9}

# "name" field is required.
$ curl -X POST localhost:9000/users -d 'weight=71.9'
map[name:This field is required]

# also support json request
$ curl -X POST -H "Content-type: application/json" localhost:9000/users -d '{"name":"bluele", "weight":71.9}'
```

## Render HTML-Form

### Simple form

```go
form := userForm()
fmt.Println(form.Html())
/* 
# Output
<input type="text" name="name"></input>
<input type="text" name="weight"></input>
*/
```

## Customize Formfield attributes

```go
customForm := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    gforms.Validators{
      gforms.Required(),
    },
    gforms.NewTextWidget(
      map[string]string{
        "class": "custom",
      },
    )),
))
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

### IntegerField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewIntegerField(
    "name",
    gforms.Validators{},
))
```

### FloatField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewFloatField(
    "name",
    gforms.Validators{},
))
```

### TextField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    gforms.Validators{},
))
```

## Support Validators

### Required validator

Added an error msg to FormInstance.Errors if the field is not provided.

```go
gforms.Validators{
  gforms.Required(),
},
```

### Regexp validator

Added an error msg to FormInstance.Errors if the regexp doesn't matched a value.

```go
gforms.Validators{
  gforms.RegexpValidator("number-\d+"),
},
```

### Email validator

Added an error msg to FormInstance.Errors if a value doesn't looks like an email address.

```go
gforms.Validators{
  gforms.EmailValidator(),
},
```

### URL Validator

Added an error msg to FormInstance.Errors if a value doesn't looks like an url.

```go
gforms.Validators{
  gforms.URLValidator(),
},
```

### MinLength Validator

Added an error msg to FormInstance.Errors if the length of value is less than specified first argument.

```go
gforms.Validators{
  gforms.MinLengthValidator(16),
},
```

### MaxLength Validator

Added an error msg to FormInstance.Errors if the length of value is greater than specified first argument.

```go
gforms.Validators{
  gforms.MaxLengthValidator(256),
},
```

### MinValueValidator

Added an error msg to FormInstance.Errors if value is less than specified first argument.

```go
gforms.Validators{
  gforms.MinValueValidator(16),
},
```

### MaxValueValidator

Added an error msg to FormInstance.Errors if value is greater than specified first argument.

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
    gforms.NewSelectWidget(
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

### RadioWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
    gforms.NewTextField(
      "lang",
      gforms.Validators{
        gforms.Required(),
      },
      gforms.NewRadioWidget(
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

### CheckboxWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
    gforms.NewTextField(
      "lang",
      gforms.Validators{
        gforms.Required(),
      },
      gforms.NewRadioWidget(
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

form := Form()
fmt.Println(form.Html())
/*
# output
<input type="checkbox" name="lang" value="0">Golang
<input type="checkbox" name="lang" value="1" disabled>Python
*/
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
