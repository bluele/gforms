package gforms

import (
	"text/template"
)

const defaultTemplates = `
{{define "TextTypeField"}}<input type="text" name="{{.Field.GetName}}" value="{{.Value | html}}"></input>{{end}}

{{define "TextWidget"}}<input type="text" name="{{.Name}}" value="{{.Value | html}}"{{range $attr, $val := .Attrs}} {{$attr}}="{{$val}}"{{end}}></input>{{end}}

{{define "SelectWidget"}}<select name="{{.Name}}"{{range $attr, $val := .Attrs}}{{$attr}}="{{$val}}"{{end}}>
{{range $idx, $val := .Options}}<option value="{{$val.Value | html}}"{{if $val.Selected }} selected{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label}}</option>
{{end}}</select>{{end}}

{{define "RadioWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="radio" name="{{$name}}" value="{{$val.Value | html}}"{{if $val.Checked }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label}}
{{end}}{{end}}

{{define "CheckboxWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="checkbox" name="{{$name}}" value="{{$val.Value | html}}"{{if $val.Checked }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label}}
{{end}}{{end}}

{{define "FileTypeField"}}<input type="file" name="{{.Field.GetName}}"></input>{{end}}
`

var Template *template.Template

func init() {
	var err error
	Template, err = template.New("gforms").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}
