package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
	"path"
	"runtime"
	"text/template"
)

type User struct {
	Email    string `gforms:"email"`
	Password string `gforms:"password"`
}

func main() {
	tpl := template.Must(template.ParseFiles(path.Join(getTemplatePath(), "post_form.html")))
	loginForm := gforms.DefineForm(
		gforms.NewFields(
			gforms.NewTextField(
				"email",
				gforms.Validators{
					gforms.Required(),
					gforms.MinLengthValidator(4),
					gforms.EmailValidator(),
				},
			),
			gforms.NewTextField(
				"password",
				gforms.Validators{
					gforms.Required(),
					gforms.MinLengthValidator(4),
					gforms.MaxLengthValidator(16),
				},
				gforms.PasswordInputWidget(map[string]string{}),
			),
		))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := loginForm(r)
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

func getTemplatePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "templates")
}
