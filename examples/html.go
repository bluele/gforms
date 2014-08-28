package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

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
				gforms.MinLengthValidator(3),
				gforms.MaxLengthValidator(32),
			},
		),
		gforms.NewFloatField(
			"weight",
			gforms.Validators{},
		),
	))

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
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
			for _, f := range form.Fields.GetList() {
				name := f.GetName()
				err, hasErr := form.Errors[name]
				if hasErr {
					fmt.Fprintf(w, "<label>%v</label><div class=\"error\">%v</div>%v\n", name, err, f.Html(form.RawData))
				} else {
					fmt.Fprintf(w, "<label>%v</label>%v\n", name, f.Html(form.RawData))
				}
			}
		}
	})
	http.ListenAndServe(":9000", nil)
}
