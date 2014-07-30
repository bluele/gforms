# GForms
A flexible forms validation and rendering library for golang web development. (under heavy development)

## Overview

* Validate HTTP request
* Rendering form-html
* Support parsing content-type "form-urlencoded", "json"

# Getting Started

## Install

```
go get github.com/bluele/gforms
```

## Define Forms

```go
var userForm gforms.Form

func initForms() {
  userForm = gforms.DefineForm(gforms.FormFields{
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(true),
        gforms.MaxLength(32),
      },
    ),
    gforms.NewFloatField(
      "weight",
      gforms.Validators{
        gforms.Required(true),
      },
    ),
  })
}
```

## Validate HTTP request

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
  form := userForm()
  if r.Method != "POST" { // Show html-form
    fmt.Fprintf(w, form.Html())
    return
  }
  err := form.ParseRequest(r)
  if err != nil { // Invalid http-request
    fmt.Fprintf(w, "%v", err)
    return
  }
  if form.IsValid() { // Validate request body
    user := User{}
    form.MapTo(&user)
    fmt.Fprintf(w, "%v", user)
  } else {
    fmt.Fprintf(w, "%v", form.Errors)
  }
}
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
var customForm gforms.Form

func initForms() {
  customForm = gforms.DefineForm(gforms.FormFields{
    gforms.NewTextField(
      "name",
      gforms.Validators{
        gforms.Required(true),
      },
      gforms.NewTextWidget(
        map[string]string{
          "class": "custom",
        },
      )),
  })
}
```

## Support Fields

### IntegerField

```go
form := gforms.DefineForm(gforms.FormFields{
  gforms.NewIntegerField(
    "name",
    nil,
})
```

### FloatField

```go
form := gforms.DefineForm(gforms.FormFields{
  gforms.NewFloatField(
    "name",
    nil,
})
```

### TextField

```go
form := gforms.DefineForm(gforms.FormFields{
  gforms.NewTextField(
    "name",
    nil,
})
```

## Support Widgets

### SelectWidget

```go
form := gforms.DefineForm(gforms.FormFields{
  gforms.NewTextField(
    "gender",
    gforms.Validators{
      gforms.Required(true),
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
})
```

### RadioWidget

```go
form := gforms.DefineForm(gforms.FormFields{
    gforms.NewTextField(
      "lang",
      gforms.Validators{
        gforms.Required(true),
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
})

```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
