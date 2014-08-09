package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
)

type Single struct {
	Check string `gforms:"check"`
}

type Multiple struct {
	Checks []string `gforms:"checks"`
}

func main() {
	SForm := gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(
			"check",
			gforms.Validators{
				gforms.Required(),
			},
			gforms.NewCheckboxWidget(
				map[string]string{},
				func() gforms.CheckboxOptions {
					return gforms.StringCheckboxOptions([][]string{
						{"Golang", "0", "false", "false"},
						{"Python", "1", "false", "true"},
					})
				},
			),
		),
	))

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		form := SForm(r)
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		if form.IsValid() {
			obj := Single{}
			form.MapTo(&obj)
			fmt.Fprintf(w, "%v <=> %v", form.CleanedData, obj)
		} else {
			fmt.Fprintf(w, "%v", form.Errors)
		}
	})

	MForm := gforms.DefineForm(gforms.NewFields(
		gforms.NewTextField(
			"checks",
			gforms.Validators{
				gforms.Required(),
			},
			gforms.NewCheckboxWidget(
				map[string]string{},
				func() gforms.CheckboxOptions {
					return gforms.StringCheckboxOptions([][]string{
						{"Golang", "0", "false", "false"},
						{"Python", "1", "false", "true"},
					})
				},
			),
		),
	))

	http.HandleFunc("/checks", func(w http.ResponseWriter, r *http.Request) {
		form := MForm(r)
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		if form.IsValid() {
			obj := Multiple{}
			form.MapTo(&obj)
			fmt.Fprintf(w, "%v <=> %v", form.CleanedData, obj)
		} else {
			fmt.Fprintf(w, "%v", form.Errors)
		}
	})

	http.ListenAndServe(":9000", nil)
}
