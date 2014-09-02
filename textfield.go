package gforms

import (
	"reflect"
)

type TextField struct {
	BaseField
}

func (f *TextField) New() FieldInterface {
	fi := new(TextFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

type TextFieldInstance struct {
	FieldInstance
}

// Create a new field for string value.
func NewTextField(name string, vs Validators, ws ...Widget) Field {
	f := new(TextField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

func (f *TextFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		v := m.rawValueAsString()
		m.Kind = reflect.String
		if v != nil {
			m.Value = *v
			m.IsNil = false
		}
	}
	return nil
}

func (f *TextFieldInstance) html() string {
	return renderTemplate("TextTypeField", newTemplateContext(f))
}

func (f *TextFieldInstance) Html() string {
	return fieldToHtml(f)
}
