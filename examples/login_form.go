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
	Name     string `gforms:"name"`
	Password string `gforms:"password"`
}

func main() {
	tpl := template.Must(template.ParseFiles(path.Join(getTemplatePath(), "login_form.html")))
	loginForm := gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
		),
		gforms.NewTextField(
			"password",
			gforms.Validators{
				gforms.Required(),
				gforms.MinLengthValidator(4),
				gforms.MaxLengthValidator(16),
			},
			gforms.NewPasswordWidget(map[string]string{}),
		),
	))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := loginForm(r)
		if r.Method != "POST" {
			err := tpl.Execute(w, form)
			if err != nil {
				panic(err)
			}
			return
		}
		if !form.IsValid() {
			err := tpl.Execute(w, form)
			if err != nil {
				panic(err)
			}
			return
		}
		user := User{}
		form.MapTo(&user)
		fmt.Fprintf(w, "%v", user)
	})
	http.ListenAndServe(":9000", nil)
}

func getTemplatePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "templates")
}
