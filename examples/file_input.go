package main

import (
	"fmt"
	"github.com/bluele/gforms"
	"mime/multipart"
	"net/http"
)

type Image struct {
	Image multipart.FileHeader `gforms:"image"`
}

func main() {
	formTpl := `<form enctype="multipart/form-data" method="post">%v<input type="submit"></form>`
	imageForm := gforms.DefineForm(gforms.NewFields(
		gforms.NewFileField(
			"image",
			gforms.Validators{
				gforms.Required(),
			},
		),
	))

	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		form := imageForm(r)
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, fmt.Sprintf(formTpl, form.Html()))
			return
		}
		if form.IsValid() {
			im := Image{}
			form.MapTo(&im)
			fmt.Fprintf(w, "File '%v' uploaded successfully.", im.Image.Filename)
		} else {
			fmt.Fprintf(w, "%v", form.Errors)
		}
	})
	http.ListenAndServe(":9000", nil)
}
