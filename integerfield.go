package gforms

import (
	"bytes"
	"errors"
	"strconv"
)

type IntegerField struct {
	BaseField
}

func (self *IntegerField) Html() string {
	if self.Widget == nil {
		return self.html()
	} else {
		return self.Widget.Html(self)
	}
}

func (self *IntegerField) html() string {
	var buffer bytes.Buffer
	Template.ExecuteTemplate(&buffer, "TextTypeField", self)
	return buffer.String()
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

func (self *IntegerField) Clean(data Data) (interface{}, error) {
	dataValue, hasField := data[self.name]
	if hasField {
		value, ok := dataValue.(string)
		if !ok {
			return nil, errors.New("Invalid type.")
		}
		if value != "" {
			v, err := strconv.Atoi(value)
			if err == nil {
				return &v, nil
			} else {
				return nil, errors.New("This field should be specified as int.")
			}
		}
	}
	return nil, nil
}
