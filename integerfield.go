package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type IntegerField struct {
	BaseField
}

func (self *IntegerField) html(vs ...string) string {
	return renderTemplate("TextTypeField", newTemplateContext(self, vs...))
}

func NewIntegerField(name string, vs Validators, ws ...Widget) *IntegerField {
	self := new(IntegerField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}

func (self *IntegerField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := m.RawValues[0]
		m.Kind = reflect.Int
		if v != "" {
			iv, err := strconv.Atoi(v)
			if err == nil {
				m.Value = iv
				m.IsNill = false
				return m, nil
			}
			return nil, errors.New("This field should be specified as int.")
		}
	}
	return nilV(), nil
}
