package gforms

import (
	"reflect"
)

type Field interface {
	Clean(Data) (*V, error)
	Validate(*V) error
	Html() string
	html() string
	GetName() string
	GetWigdet() Widget
}

type ValidationError interface {
	Error() string
}

type BaseField struct {
	name       string
	validators Validators
	Widget     Widget
	Field
}

func (self *BaseField) GetName() string {
	return self.name
}

func (self *BaseField) GetWigdet() Widget {
	return self.Widget
}

func (self *BaseField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := m.RawValues[0]
		m.Kind = reflect.String
		if v != "" {
			m.Value = v
			m.IsNill = false
			return m, nil
		}
	}
	return nilV(), nil
}

func (self *BaseField) Validate(value *V) error {
	if self.validators == nil {
		return nil
	}
	for _, v := range self.validators {
		err := v.Validate(value)
		if err != nil {
			return err
		}
	}
	return nil
}
