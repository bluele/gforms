package gforms

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Validator interface {
	Name() string
	Validate(*FieldInstance, *FormInstance) error
}

type Validators []Validator

type required struct {
	Message string
	Validator
}

// Returns error if the field is not provided.
func Required(message ...string) required {
	vl := required{}
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = "This field is required."
	}
	return vl
}

func (vl required) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || (v.Kind == reflect.String && v.Value == "") {
		return errors.New(vl.Message)
	}
	return nil
}

type maxLengthValidator struct {
	Length  int
	Message string
	Validator
}

// Returns error if the length of value is greater than length argument.
func MaxLengthValidator(length int, message ...string) maxLengthValidator {
	vl := maxLengthValidator{}
	vl.Length = length
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = fmt.Sprintf("Ensure this value has at most %v characters.", vl.Length)
	}
	return vl
}

func (vl maxLengthValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.String || v.Value == "" {
		return nil
	}
	s := v.Value.(string)
	if len(s) > vl.Length {
		return errors.New(vl.Message)
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
	vl := minLengthValidator{}
	vl.Length = length
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = fmt.Sprintf("Ensure this value has at least %v characters", vl.Length)
	}
	return vl
}

func (vl minLengthValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.String || v.Value == "" {
		return nil
	}
	s := v.Value.(string)
	if len(s) < vl.Length {
		return errors.New(vl.Message)
	}
	return nil
}

type maxValueValidator struct {
	Value   int
	Message string
	Validator
}

func MaxValueValidator(value int, message ...string) maxValueValidator {
	vl := maxValueValidator{}
	vl.Value = value
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = fmt.Sprintf("Ensure this value is less than or equal to %v.", vl.Value)
	}
	return vl
}

func (vl maxValueValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.Int {
		return nil
	}
	iv := v.Value.(int)
	if iv > vl.Value {
		return errors.New(vl.Message)
	}
	return nil
}

type minValueValidator struct {
	Value   int
	Message string
	Validator
}

func MinValueValidator(value int, message ...string) minValueValidator {
	vl := minValueValidator{}
	vl.Value = value
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = fmt.Sprintf("Ensure this value is greater than or equal to %v.", vl.Value)
	}
	return vl
}

func (vl minValueValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.Int {
		return nil
	}
	iv := v.Value.(int)
	if iv < vl.Value {
		return errors.New(vl.Message)
	}
	return nil
}

type regexpValidator struct {
	re      *regexp.Regexp
	Message string
	Validator
}

func (vl regexpValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.String || v.Value == "" {
		return nil
	}
	sv := v.Value.(string)
	if !vl.re.MatchString(sv) {
		return errors.New(vl.Message)
	}
	return nil
}

// The regular expression pattern to search for the provided value.
// Returns error if regxp#MatchString is False.
func RegexpValidator(regex string, message ...string) regexpValidator {
	vl := regexpValidator{}
	re, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	vl.re = re
	if len(message) > 0 {
		vl.Message = message[0]
	} else {
		vl.Message = fmt.Sprintf("Enter a valid value.")
	}
	return vl
}

// An EmailValidator that ensures a value looks like an international email address.
func EmailValidator(message ...string) regexpValidator {
    regex := `^.+@.+$` // international email can include UTF8 characters.  Better to have false positives than false negatives.
    if len(message) > 0 {
        return RegexpValidator(regex, message[0])
    } else {
        return RegexpValidator(regex, "Enter a valid email address.")
    }
}

// A FullNameValidator that ensures that we have a full name (e.g. 'John Doe').
func FullNameValidator(message ...string) regexpValidator {
    regex := `^[\p{L}]+( [\p{L}]+)+$`
    if len(message) > 0 {
        return RegexpValidator(regex, message[0])
    } else {
        return RegexpValidator(regex, "Enter a valid full name (i.e. 'John Doe').")
    }
}

// An URLValidator that ensures a value looks like an url.
func URLValidator(message ...string) regexpValidator {
    regex := `^(https?|ftp)(:\/\/[-_.!~*\'()a-zA-Z0-9;\/?:\@&=+\$,%#]+)$`
    if len(message) > 0 {
        return RegexpValidator(regex, message[0])
    } else {
        return RegexpValidator(regex, "Enter a valid url.")
    }
}

type passwordStrengthValidator struct {
	RequiredStrength int
    Message string
    Validator
}
// A PasswordStrengthValidator that ensures that the password is complex enough.
// requiredStrength:
//	0: Horrible
//	1: Weak
//	2: Medium
//	3: Strong
//	4: Very Strong		
func PasswordStrengthValidator(requiredStrength int, message ...string) passwordStrengthValidator {
    vl := passwordStrengthValidator{}
    vl.RequiredStrength = requiredStrength
    if len(message) > 0 {
        vl.Message = message[0]
    } else {
        vl.Message = "Password isn't strong enough.  Add a mix of uppers, lowers, numbers, and special characters."
    }
    return vl
}
func (vl passwordStrengthValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
    v := fi.V
    if v.IsNil || v.Kind != reflect.String {
        return nil
    }
    sv := v.Value.(string)
    
    password := New(sv)
    password.ProcessPassword()
    if password.CommonPassword {
		return errors.New("Your password is a common password.  Try making it harder to guess.")
	}
    if password.Score < vl.RequiredStrength {
        return errors.New(vl.Message)
    }
    return nil
}

type fieldMatchValidator struct {
	FieldMatchName string
    Message string
    Validator
}
// A FieldMatchValidator ensures that this field matches the field [FieldMatchName].
func FieldMatchValidator(fieldMatchName string, message ...string) fieldMatchValidator {
	vl := fieldMatchValidator{}
	vl.FieldMatchName = fieldMatchName
	if len(message) > 0 {
        vl.Message = message[0]
    } else {
        vl.Message = fieldMatchName + " fields must match"
    }
    return vl
}
func (vl fieldMatchValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
    v := fi.V
    if v.IsNil || v.Kind != reflect.String {
        return nil
    }
    sv := v.Value.(string)
    
    if sv != fo.Data[vl.FieldMatchName].RawStr {
        return errors.New(vl.Message)
    }
    return nil
}

type fnValidator struct {
	ValidationFn func(fi *FieldInstance, fo *FormInstance) error
    Validator
}
// A FieldMatchValidator ensures that this field matches the field [FieldMatchName].
func FnValidator(validationFn func(fi *FieldInstance, fo *FormInstance) error) fnValidator {
	vl := fnValidator{}
	vl.ValidationFn = validationFn
    return vl
}
func (vl fnValidator) Validate(fi *FieldInstance, fo *FormInstance) error {
    return vl.ValidationFn(fi, fo)
}
