package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

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

type User struct {
	Name   string  `gforms:"name"`
	Weight float32 `gforms:"weight"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	form := userForm()
	if r.Method != "POST" {
		fmt.Fprintf(w, form.Html())
		return
	}
	err := form.ParseRequest(r)
	if err != nil { // Invalid http request
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

func main() {
	initForms()
	http.HandleFunc("/users", createUserHandler)
	http.ListenAndServe(":9000", nil)
}
