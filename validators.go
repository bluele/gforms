package gforms

import (
	"errors"
	"fmt"
)

type Validator interface {
	Name() string
	Validate(*V) error
}

type Validators []Validator

type required struct {
	IsRequired bool
	Message    string
	Validator
}

func (self required) Validate(value *V) error {
	if self.IsRequired && value.IsNill {
		return errors.New(self.Message)
	}
	return nil
}

func Required(isRequired bool, message ...string) required {
	self := new(required)
	self.IsRequired = isRequired
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = "This field is required"
	}
	return *self
}

type maxLength struct {
	Length  int
	Message string
	Validator
}

func MaxLength(length int, message ...string) maxLength {
	self := new(maxLength)
	self.Length = length
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = fmt.Sprintf("This field cannot be longer than %v characters.", self.Length)
	}
	return *self
}

func (self maxLength) Validate(value *V) error {
	s, ok := value.Value.(*string)
	if ok && len(*s) > self.Length {
		return errors.New(self.Message)
	}
	return nil
}

type minLength struct {
	Length  int
	Message string
	Validator
}

func MinLength(length int, message ...string) minLength {
	self := new(minLength)
	self.Length = length
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = fmt.Sprintf("This field cannot be shorter than %v characters.", self.Length)
	}
	return *self
}

func (self minLength) Validate(value *V) error {
	s, ok := value.Value.(*string)
	if ok && len(*s) < self.Length {
		return errors.New(self.Message)
	}
	return nil
}
