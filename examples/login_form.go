package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

type User struct {
	Name     string `gforms:"name"`
	Password string `gforms:"password"`
}

func main() {
	userForm := gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
		),
		gforms.NewTextField(
			"password",
			gforms.Validators{
				gforms.MinLengthValidator(4),
				gforms.MaxLengthValidator(16),
			},
			gforms.NewPasswordWidget(map[string]string{}),
		),
	))

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := userForm(r)
		if r.Method != "POST" {
			html := `<form method="post">`
			for _, f := range form.Fields.GetList() {
				html += fmt.Sprintf(`<label>%v</label>`, f.GetName())
				html += f.Html() + "<br />"
			}
			html += `<input type="submit"></form>`
			fmt.Fprintf(w, html)
			return
		}
		if form.IsValid() { // Validate request body
			user := User{}
			form.MapTo(&user)
			fmt.Fprintf(w, "%v", user)
		} else {
			html := `<form method="post">`
			for _, f := range form.Fields.GetList() {
				name := f.GetName()
				err, hasErr := form.Errors[name]
				if hasErr {
					html += fmt.Sprintf("<label>%v</label><div class=\"error\">%v</div>%v<br />", name, err, f.Html(form.RawData))
				} else {
					html += fmt.Sprintf("<label>%v</label>%v<br />", name, f.Html(form.RawData))
				}
			}
			html += `<input type="submit"></form>`
			fmt.Fprintf(w, html)
		}
	})
	http.ListenAndServe(":9000", nil)
}
