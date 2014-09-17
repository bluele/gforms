package gforms

import (
	"errors"
	"reflect"
	"time"
)

var DefaultDateFormat string = "2006-01-02"

var DefaultDateTimeFormat string = "2006-01-02 15:04:05"

// It maps value to FormInstance.CleanedData as type `time.Time`.
type DateTimeField struct {
	BaseField
	Format       string
	ErrorMessage string
}

// Create a new DateField instance.
func (f *DateTimeField) New() FieldInterface {
	fi := new(DateTimeFieldInstance)
	fi.Format = f.Format

	if f.ErrorMessage == "" {
		fi.ErrorMessage = "This field should be specified as date format."
	} else {
		fi.ErrorMessage = f.ErrorMessage
	}
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for DateTimeField
type DateTimeFieldInstance struct {
	FieldInstance
	Format       string
	ErrorMessage string
}

// Create a new DateTimeField with validators and widgets.
func NewDateTimeField(name string, format string, vs Validators, ws ...Widget) *DateTimeField {
	f := new(DateTimeField)
	f.name = name
	f.validators = vs
	f.Format = format
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

// Get a value from request data, and clean it as type string
func (f *DateTimeFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		v := m.rawValueAsString()
		m.Kind = reflect.Struct
		if v != nil {
			t, err := time.Parse(f.Format, *v)
			if err != nil {
				return errors.New(f.ErrorMessage)
			}
			m.Value = t
			m.IsNil = false
		}
	}
	return nil
}

func (f *DateTimeFieldInstance) html() string {
	return renderTemplate("TextTypeField", newTemplateContext(f))
}

func (f *DateTimeFieldInstance) Html() string {
	return fieldToHtml(f)
}
