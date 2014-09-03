package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
	"path"
	"runtime"
	"text/template"
)

type Lang struct {
	Name string `gforms:"name"`
}

func main() {
	tpl := template.Must(template.ParseFiles(path.Join(getTemplatePath(), "post_form.html")))
	langForm := gforms.DefineModelForm(Lang{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
			gforms.SelectWidget(
				map[string]string{},
				func() gforms.SelectOptions {
					return gforms.StringSelectOptions([][]string{
						{"Select...", "", "true", "false"},
						{"Golang", "golang", "false", "false"},
						{"Python", "python", "false", "false"},
						{"C", "c", "false", "true"},
					})
				},
			),
		),
	))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := langForm(r)
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
		lang := form.GetModel().(Lang)
		fmt.Fprintf(w, "ok: %v", lang)
	})

	http.ListenAndServe(":9000", nil)
}

func getTemplatePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "templates")
}
