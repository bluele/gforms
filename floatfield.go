package gforms

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

type FloatField struct {
	BaseField
}

func (self *FloatField) Html() string {
	if self.Widget == nil {
		return self.html()
	} else {
		return self.Widget.Html(self)
	}
}

func (self *FloatField) html() string {
	var buffer bytes.Buffer
	Template.ExecuteTemplate(&buffer, "TextTypeField", self)
	return buffer.String()
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
				m.IsNill = false
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
