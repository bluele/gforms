package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

type Lang struct {
	Name string `gforms:"name"`
}

func main() {
	formTpl := `<form method="post">%v<input type="submit"></form>`
	langForm := gforms.DefineModelForm(Lang{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
			},
			gforms.RadioSelectWidget(
				map[string]string{},
				func() gforms.RadioOptions {
					return gforms.StringRadioOptions([][]string{
						{"Golang", "golang", "false", "false"},
						{"Python", "python", "false", "false"},
						{"C", "c", "false", "true"},
					})
				},
			),
		),
	))
	http.HandleFunc("/lang", func(w http.ResponseWriter, r *http.Request) {
		form := langForm(r)
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, fmt.Sprintf(formTpl, form.Html()))
			return
		}
		if form.IsValid() {
			obj := form.GetModel()
			fmt.Fprintf(w, "%v <=> %v", form.CleanedData, obj)
		} else {
			fmt.Fprintf(w, "%v", form.Errors())
		}
	})

	http.ListenAndServe(":9000", nil)
}
