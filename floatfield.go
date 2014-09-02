package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type FloatField struct {
	BaseField
	ErrorMessage string
}

func (f *FloatField) New() FieldInterface {
	fi := new(FloatFieldInstance)
	if f.ErrorMessage == "" {
		fi.ErrorMessage = "This field should be specified as float."
	} else {
		fi.ErrorMessage = f.ErrorMessage
	}
	fi.Model = f
	fi.V = nilV("")
	return fi
}

type FloatFieldInstance struct {
	FieldInstance
	ErrorMessage string
}

func NewFloatField(name string, vs Validators, ws ...Widget) *FloatField {
	f := new(FloatField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

func (f *FloatFieldInstance) Clean(data Data) error {
	m, hasField := data[f.GetName()]
	if hasField {
		f.V = m
		v := m.rawValueAsString()
		m.Kind = reflect.Float64
		if v != nil && (*v) != "" {
			fv, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				m.Value = fv
				m.IsNil = false
				return nil
			}
			return errors.New(f.ErrorMessage)
		}
	}
	return nil
}

func (f *FloatFieldInstance) html() string {
	return renderTemplate("TextTypeField", newTemplateContext(f))
}

func (f *FloatFieldInstance) Html() string {
	return fieldToHtml(f)
}
