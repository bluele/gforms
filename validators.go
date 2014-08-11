package gforms

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Validator interface {
	Name() string
	Validate(*V) error
}

type Validators []Validator

type required struct {
	Message string
	Validator
}

func (self required) Validate(value *V) error {
	if value.IsNil {
		return errors.New(self.Message)
	}
	return nil
}

func Required(message ...string) required {
	self := new(required)
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
		self.Message = fmt.Sprintf("Ensure this value has at most %v characters.", self.Length)
	}
	return *self
}

func (self maxLength) Validate(value *V) error {
	if value.IsNil || value.Kind != reflect.String {
		return nil
	}
	s := value.Value.(string)
	if len(s) > self.Length {
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
		self.Message = fmt.Sprintf("Ensure this value has at least %v characters", self.Length)
	}
	return *self
}

func (self minLength) Validate(value *V) error {
	if value.IsNil || value.Kind != reflect.String {
		return nil
	}
	s := value.Value.(string)
	if len(s) < self.Length {
		return errors.New(self.Message)
	}
	return nil
}

type regexpValidator struct {
	re      *regexp.Regexp
	Message string
	Validator
}

func (self regexpValidator) Validate(value *V) error {
	if value.IsNil || value.Kind != reflect.String {
		return nil
	}
	s := value.Value.(string)
	if !self.re.MatchString(s) {
		return errors.New(self.Message)
	}
	return nil
}

func RegexpValidator(regex string, message ...string) regexpValidator {
	self := regexpValidator{}
	re, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	self.re = re
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = fmt.Sprintf("Enter a valid value.")
	}
	return self
}

func EmailValidator(message ...string) regexpValidator {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if len(message) > 0 {
		return RegexpValidator(regex, message[0])
	} else {
		return RegexpValidator(regex, "Enter a valid email address.")
	}
}
