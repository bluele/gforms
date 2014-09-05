package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
	"text/template"
)

var tplText string = `
<form method="get">
{{range $i, $field := .Fields}}
  <label>{{$field.GetName}}: </label>{{$field.Html}}
  {{range $ei, $err := $field.Errors}}
  <label class="error">{{$err}}</label>
  {{end}}
  <br />
{{end}}
<input type="submit">
</form>
`

func main() {
	tpl, _ := template.New("tpl").Parse(tplText)
	searchForm := gforms.DefineForm(
		gforms.NewFields(
			gforms.NewTextField(
				"query",
				gforms.Validators{
					gforms.Required(),
					gforms.MinLengthValidator(2),
				},
			),
		))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := searchForm.FromUrlValues(r.URL.Query())
		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
		fmt.Fprintf(w, "input query: %v", form.CleanedData["query"])
	})
	http.ListenAndServe(":9000", nil)
}
