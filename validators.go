package gforms

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Validator interface {
	Name() string
	Validate(*V, CleanedData) error
}

type Validators []Validator

type required struct {
	Message string
	Validator
}

func (self required) Validate(value *V, cleanedData CleanedData) error {
	if value.IsNil {
		return errors.New(self.Message)
	}
	return nil
}

// Returns error if the field is not provided.
func Required(message ...string) required {
	self := required{}
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = "This field is required"
	}
	return self
}

type maxLengthValidator struct {
	Length  int
	Message string
	Validator
}

// Returns error if the length of value is greater than length argument.
func MaxLengthValidator(length int, message ...string) maxLengthValidator {
	self := maxLengthValidator{}
	self.Length = length
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = fmt.Sprintf("Ensure this value has at most %v characters.", self.Length)
	}
	return self
}

func (self maxLengthValidator) Validate(value *V, cleanedData CleanedData) error {
	if value.IsNil || value.Kind != reflect.String {
		return nil
	}
	s := value.Value.(string)
	if len(s) > self.Length {
		return errors.New(self.Message)
	}
	return nil
}

type minLengthValidator struct {
	Length  int
	Message string
	Validator
}

// Returns error if the length of value is less than length argument.
func MinLengthValidator(length int, message ...string) minLengthValidator {
	self := minLengthValidator{}
	self.Length = length
	if len(message) > 0 {
		self.Message = message[0]
	} else {
		self.Message = fmt.Sprintf("Ensure this value has at least %v characters", self.Length)
	}
	return self
}

func (self minLengthValidator) Validate(value *V, cleanedData CleanedData) error {
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

func (self regexpValidator) Validate(value *V, cleanedData CleanedData) error {
	if value.IsNil || value.Kind != reflect.String {
		return nil
	}
	s := value.Value.(string)
	if !self.re.MatchString(s) {
		return errors.New(self.Message)
	}
	return nil
}

// The regular expression pattern to search for the provided value.
// Returns error if regxp#MatchString is False.
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

// An EmailValidator that ensures a value looks like an email address.
func EmailValidator(message ...string) regexpValidator {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if len(message) > 0 {
		return RegexpValidator(regex, message[0])
	} else {
		return RegexpValidator(regex, "Enter a valid email address.")
	}
}
