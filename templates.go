package gforms

import (
	"text/template"
)

const defaultTemplates = `
{{define "TextTypeField"}}<input type="text" name="{{.Field.GetName | html}}" value="{{.Value | html}}"></input>{{end}}

{{define "SimpleWidget"}}<input type="{{.Type | html}}" name="{{.Name}}" value="{{.Value | html}}"{{range $attr, $val := .Attrs}} {{$attr}}="{{$val}}"{{end}}></input>{{end}}

{{define "SelectWidget"}}<select name="{{.Name | html}}"{{range $attr, $val := .Attrs}}{{$attr | html}}="{{$val | html}}"{{end}}>
{{range $idx, $val := .Options}}<option value="{{$val.Value | html}}"{{if $val.Selected }} selected{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}</option>
{{end}}</select>{{end}}

{{define "RadioWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="radio" name="{{$name | html}}" value="{{$val.Value | html}}"{{if $val.Checked}} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}

{{define "CheckboxWidget"}}{{ $name := .Name}}{{range $idx, $val := .Options}}<input type="checkbox" name="{{$name | html}}" value="{{$val.Value | html}}"{{if $val.Checked}} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}

{{define "FileTypeField"}}<input type="file" name="{{.Field.GetName | html}}"></input>{{end}}
`

var Template *template.Template

func init() {
	var err error
	Template, err = template.New("gforms").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}
