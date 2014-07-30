package gforms

import (
	"errors"
)

type Field interface {
	Clean(data Data) (interface{}, error)
	Validate(value interface{}) error
	Html() string
	html() string
	GetName() string
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

func (self *BaseField) Clean(data Data) (interface{}, error) {
	dataValue, hasField := data[self.name]
	if hasField {
		value, ok := dataValue.(string)
		if !ok {
			return nil, errors.New("Invalid type.")
		}
		if value != "" {
			return &value, nil
		}
	}
	return nil, nil
}

func (self *BaseField) Validate(value interface{}) error {
	if self.Widget != nil {
		err := self.Widget.Validate(value)
		if err != nil {
			return err
		}
	}
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
