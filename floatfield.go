package gforms

import (
	"bytes"
	"errors"
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

func (self *FloatField) Clean(data Data) (interface{}, error) {
	dataValue, hasField := data[self.name]
	if hasField {
		value, ok := dataValue.(string)
		if !ok {
			return nil, errors.New("Invalid type.")
		}
		if value != "" {
			v, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return &v, nil
			} else {
				return nil, errors.New("This field should be specified as float.")
			}
		}
	}
	return nil, nil
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
