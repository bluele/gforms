package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

type User struct {
	Name   string  `gforms:"name"`
	Age    int     `gforms:"age"`
	Weight float64 `gforms:"weight"`
	Ext    string  `gforms:"-"`
}

func main() {
	Form := gforms.DefineModelForm(User{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
		),
	))
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		form := Form(r)
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		if form.IsValid() {
			user := form.GetModel().(User)
			fmt.Fprintf(w, "%v => %v", form.CleanedData, user)
		} else {
			fmt.Println(form.Html())
			fmt.Fprintf(w, "%v", form.Errors())
		}
	})
	http.ListenAndServe(":9000", nil)
}
