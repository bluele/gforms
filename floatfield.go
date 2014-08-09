package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type FloatField struct {
	BaseField
}

func (self *FloatField) Html(rd RawData) string {
	return fieldToHtml(self, rd)
}

func (self *FloatField) html(vs ...string) string {
	return renderTemplate("TextTypeField", newTemplateContext(self, vs...))
}

func (self *FloatField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := m.RawValues[0]
		m.Kind = reflect.Float64
		if v != "" {
			fv, err := strconv.ParseFloat(v, 64)
			if err == nil {
				m.Value = fv
				m.IsNil = false
				return m, nil
			}
			return nil, errors.New("This field should be specified as float.")
		}
	}
	return nilV(), nil
}

func NewFloatField(name string, vs Validators, ws ...Widget) *FloatField {
	self := new(FloatField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}
