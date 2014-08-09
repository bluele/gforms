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

func (self CustomValidator) Validate(v *gforms.V) error {
	if !v.IsNil && v.Kind == reflect.String {
		s := v.Value.(string)
		for _, t := range self.Langs {
			if s == t {
				return nil
			}
		}
		return errors.New(fmt.Sprintf("Unknown lang: %v", s))
	}
	return nil
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

	http.HandleFunc("/langs", func(w http.ResponseWriter, r *http.Request) {
		form := langForm(r)
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		if form.IsValid() { // Validate request body
			lang := form.GetModel()
			fmt.Fprintf(w, "%v", lang)
		} else {
			fmt.Fprintf(w, "%v", form.Errors)
		}
	})
	http.ListenAndServe(":9000", nil)
}
