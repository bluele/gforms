package gforms

import (
	"bytes"
	"reflect"
)

// It maps value to FormInstance.CleanedData as type `bool`.
type NullBooleanField struct {
	BaseField
}

// Create a new NullBooleanField instance.
func (f *NullBooleanField) New() FieldInterface {
	fi := new(NullBooleanFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for NullBooleanField.
type NullBooleanFieldInstance struct {
	FieldInstance
}

type nullBooleanContext struct {
	Field   FieldInterface
	Checked bool
}

// Create a new NullBooleanField with validators and widgets.
func NewNullBooleanField(name string, vs Validators, ws ...Widget) *NullBooleanField {
	f := new(NullBooleanField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

// Get a value from request data, and clean it as type `bool`.
func (f *NullBooleanFieldInstance) Clean(data Data) error {
	m, hasField := data[f.GetName()]
	if hasField {
		f.V = m
		v := false
		if m.Kind == reflect.String {
			vs := m.rawValueAsString()
			if vs != nil {
				v = true
			}
		} else if m.Kind == reflect.Bool {
			v = m.rawValueAsBool()
		}
		m.Value = v
		m.Kind = reflect.Bool
		m.IsNil = false
		return nil
	}
	nv := nilV("")
	f.V = nv
	return nil
}

func (f *NullBooleanFieldInstance) html() string {
	var buffer bytes.Buffer
	cx := new(nullBooleanContext)
	cx.Field = f
	checked, _ := f.V.Value.(bool)
	cx.Checked = checked
	err := Template.ExecuteTemplate(&buffer, "BooleanTypeField", cx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Get as HTML format.
func (f *NullBooleanFieldInstance) Html() string {
	return fieldToHtml(f)
}
