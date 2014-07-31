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
	SForm := gforms.DefineForm(gforms.FormFields{
		gforms.NewTextField(
			"check",
			gforms.Validators{
				gforms.Required(true),
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
	})

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		form := SForm()
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		err := form.ParseRequest(r)
		if err != nil { // Invalid http request
			fmt.Fprintf(w, "%v", err)
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

	MForm := gforms.DefineForm(gforms.FormFields{
		gforms.NewTextField(
			"checks",
			gforms.Validators{
				gforms.Required(true),
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
	})

	http.HandleFunc("/checks", func(w http.ResponseWriter, r *http.Request) {
		form := MForm()
		if r.Method != "POST" {
			fmt.Fprintf(w, form.Html())
			return
		}
		err := form.ParseRequest(r)
		if err != nil { // Invalid http request
			fmt.Fprintf(w, "%v", err)
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
