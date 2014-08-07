package gforms

import (
	"text/template"
)

const defaultTemplates = `
{{define "TextTypeField"}}<input type="text" name="{{.Field.GetName}}" value="{{.Value}}"></input>{{end}}

{{define "TextWidget"}}<input type="text" name="{{.Name}}" value=""{{range $attr, $val := .Attrs}} {{$attr}}="{{$val}}"{{end}}></input>{{end}}

{{define "SelectWidget"}}<select {{range $attr, $val := .Attrs}}{{$attr}}="{{$val}}"{{end}}>
{{range $idx, $val := .Options}}<option value="{{$val.Value}}">{{$val.Label}}</option>
{{end}}</select>{{end}}

{{define "RadioWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="radio" name="{{$name}}" value="{{$val.Value}}"{{if $val.Checked }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label}}
{{end}}{{end}}

{{define "CheckboxWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="checkbox" name="{{$name}}" value="{{$val.Value}}"{{if $val.Checked }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label}}
{{end}}{{end}}
`

var Template *template.Template

func init() {
	var err error
	Template, err = template.New("gforms").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}
