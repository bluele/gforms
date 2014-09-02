package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type FloatField struct {
	BaseField
}

func (f *FloatField) New() FieldInterface {
	fi := new(FloatFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

type FloatFieldInstance struct {
	FieldInstance
}

func NewFloatField(name string, vs Validators, ws ...Widget) Field {
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
			return errors.New("This field should be specified as float.")
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
