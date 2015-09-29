package gforms

import (
	"bytes"
	"text/template"
)

const defaultTemplates = `
{{define "TextTypeField"}}<input type="text" name="{{.Field.GetName | html}}" value="{{.Value | html}}"></input>{{end}}
{{define "BooleanTypeField"}}<input type="checkbox" name="{{.Field.GetName | html}}"{{if .Checked}} checked{{end}}>{{end}}
{{define "SimpleWidget"}}<input type="{{.Type | html}}" name="{{.Field.GetName | html}}" value="{{.Value | html}}"{{range $attr, $val := .Attrs}} {{$attr | html}}="{{$val | html}}"{{end}}></input>{{end}}
{{define "SelectWidget"}}<select {{if .Multiple }}multiple {{end}}name="{{.Field.GetName | html}}"{{range $attr, $val := .Attrs}}{{$attr | html}}="{{$val | html}}"{{end}}>
{{range $idx, $val := .Options}}<option value="{{$val.Value | html}}"{{if $val.Selected }} selected{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}</option>
{{end}}</select>{{end}}
{{define "RadioWidget"}}{{$name := .Field.GetName}}{{range $idx, $val := .Options}}<input type="radio" name="{{$name | html}}" value="{{$val.Value | html}}"{{if or $val.Checked (eq $.Field.GetV.RawStr $val.Value) }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}
{{define "CheckboxMultipleWidget"}}{{$name := .Field.GetName}}{{range $idx, $val := .Options}}<input type="checkbox" name="{{$name | html}}" value="{{$val.Value | html}}"{{if $val.Checked}} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}
`

// all templates of Field and Widget
var Template *template.Template

func init() {
	var err error
	Template, err = template.New("gforms").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}

type templateContext struct {
	Field FieldInterface
	Value string
}

func newTemplateContext(f FieldInterface) templateContext {
	ctx := templateContext{
		Field: f,
	}
	v := f.GetV()
	if v != nil {
		ctx.Value = v.RawStr
	}
	return ctx
}

func renderTemplate(name string, ctx interface{}) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, name, ctx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
