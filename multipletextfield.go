package gforms

import (
	"reflect"
)

// It maps value to FormInstance.CleanedData as type `[]string`.
type MultipleTextField struct {
	BaseField
}

// Create a new MultipleField instance.
func (f *MultipleTextField) New() FieldInterface {
	fi := new(MultipleTextFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for MultipleTextField.
type MultipleTextFieldInstance struct {
	FieldInstance
}

// Create a new MultipleTextField with validators and widgets.
func NewMultipleTextField(name string, vs Validators, ws ...Widget) *MultipleTextField {
	f := new(MultipleTextField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	} else {
		f.widget = SelectMultipleWidget(map[string]string{}, nil)
	}
	return f
}

// Get a value from request data, and clean it as type `[]string`.
func (f *MultipleTextFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		m.Kind = reflect.Slice
		m.Value = m.rawValueAsStringArray()
		m.IsNil = false
		return nil
	}
	return nil
}

func (f *MultipleTextFieldInstance) html() string {
	return ""
}

// Get as HTML format.
func (f *MultipleTextFieldInstance) Html() string {
	return fieldToHtml(f)
}
