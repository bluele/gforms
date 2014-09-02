package main

import (
	"errors"
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
	"reflect"
)

type Lang struct {
	Name string `gforms:"name"`
}

type CustomValidator struct {
	Langs []string
	gforms.Validator
}

func (vl CustomValidator) Validate(fi *gforms.FieldInstance, fo *gforms.FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.String || v.Value == "" {
		return nil
	}
	for _, t := range vl.Langs {
		if v.Value == t {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unknown lang: %v", v.Value))
}

func main() {
	langForm := gforms.DefineModelForm(Lang{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
				CustomValidator{
					Langs: []string{"golang", "python", "c"},
				},
			},
		),
	))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := langForm(r)
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		if form.IsValid() { // Validate request body
			lang := form.GetModel().(Lang)
			fmt.Fprintf(w, "%v", lang)
		} else {
			fmt.Fprintf(w, "%v", form.Errors())
		}
	})
	http.ListenAndServe(":9000", nil)
}
