package gforms

import (
	"reflect"
)

// It maps value to FormInstance.CleanedData as type `string`.
type TextField struct {
	BaseField
}

// Create a new TextField instance.
func (f *TextField) New() FieldInterface {
	fi := new(TextFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for TextField
type TextFieldInstance struct {
	FieldInstance
}

// Create a new TextField with validators and widgets.
func NewTextField(name string, vs Validators, ws ...Widget) *TextField {
	f := new(TextField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

// Get a value from request data, and clean it as type `string`.
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

// Get as HTML format.
func (f *TextFieldInstance) Html() string {
	return fieldToHtml(f)
}
