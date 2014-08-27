package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

type User struct {
	Name   string  `gforms:"name"`
	Email  string  `gforms:"email"`
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
		gforms.NewTextField(
			"email",
			gforms.Validators{
				gforms.Required(),
				gforms.EmailValidator(),
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
			w.Header().Set("Content-Type", "text/html")
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

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		form := userForm.FromValues(map[string][]string{
			"name":   {"bluele"},
			"email":  {"junkxdev@gmail.com"},
			"weight": {"71.9"},
		})
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
