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

## Usage

### Define Form

```go
userForm := gforms.DefineForm(gforms.NewFields(
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
      gforms.Required(false),
    },
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
        gforms.Required(true),
        gforms.MaxLength(32),
      },
    ),
    gforms.NewFloatField(
      "weight",
      gforms.Validators{
        gforms.Required(false),
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

```

### Define Form by struct model

```go
type User struct {
  Name   string  `gforms:"name"`
  Weight float32 `gforms:"weight"`
}

func initForm() {
  userForm := gforms.DefineModelForm(
    User{},
    gforms.NewFields()
  )
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
  userForm := gforms.DefineForm(User{})

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
      gforms.Required(true),
    },
    gforms.NewTextWidget(
      map[string]string{
        "class": "custom",
      },
    )),
))
```

## Support Fields

### IntegerField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewIntegerField(
    "name",
    nil,
))
```

### FloatField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewFloatField(
    "name",
    nil,
))
```

### TextField

```go
form := gforms.DefineForm(gforms.NewFields(
  gforms.NewTextField(
    "name",
    nil,
))
```

## Support Widgets

### SelectWidget

```go
Form := gforms.DefineForm(gforms.NewFields(
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
))

form = Form()
fmt.Println(form.Html())
/*
# output
<input type="radio" name="lang" value="0">Golang
<input type="radio" name="lang" value="1" disabled>Python
*/
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
