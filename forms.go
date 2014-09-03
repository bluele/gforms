package gforms

import (
	"bytes"
	"net/http"
	"reflect"
)

type Form func(...*http.Request) *FormInstance

// cleaned data for all fields.
type CleanedData map[string]interface{}

type FormInstance struct {
	fieldInstances *FieldInterfaces
	Data           Data
	CleanedData    CleanedData
	ParseError     error
}

func (f *FormInstance) GetField(name string) (FieldInterface, bool) {
	v, ok := f.fieldInstances.nameMap[name]
	return v, ok
}

func (f *FormInstance) Fields() []FieldInterface {
	return f.fieldInstances.list
}

func (f *FormInstance) Errors() Errors {
	errs := map[string][]string{}
	var err []string
	for _, field := range f.fieldInstances.list {
		name := field.GetModel().GetName()
		err = field.Errors()
		if err != nil && len(err) > 0 {
			errs[name] = err
		}
	}
	return errs
}

func (f *FormInstance) IsValid() bool {
	isValid := true
	f.CleanedData = CleanedData{}

	for _, field := range f.fieldInstances.list {
		var err error
		name := field.GetModel().GetName()
		err = field.Clean(f.Data)
		if err != nil {
			field.SetErrors([]string{err.Error()})
			isValid = false
			continue
		}

		errs := field.Validate(f)
		if len(errs) > 0 {
			field.SetErrors(errs)
			isValid = false
			continue
		}

		if !field.GetV().IsNil {
			f.CleanedData[name] = field.GetV().Value
		}
	}
	return isValid
}

func (f *FormInstance) parseRequest(req *http.Request) error {
	data, err := parseReuqestBody(req)
	if err != nil {
		return err
	}
	if data == nil {
		return nil
	}
	f.Data = *data
	return nil
}

func (f *FormInstance) Html() string {
	var html bytes.Buffer
	for _, field := range f.fieldInstances.list {
		html.WriteString(field.Html() + "\r\n")
	}
	return html.String()
}

func DefineForm(fs *Fields) Form {
	return func(r ...*http.Request) *FormInstance {
		f := new(FormInstance)
		f.fieldInstances = newFieldInterfaces(fs)
		if len(r) > 0 {
			f.ParseError = f.parseRequest(r[0])
		}
		return f
	}
}

// maps cleanedData to specified model.
func (fi *FormInstance) MapTo(model interface{}) {
	if fi.CleanedData == nil {
		panic("MapTo method should be called after calling IsValid() method.")
	}
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		panic("Argument should be specified pointer type.")
	}
	mType := reflect.TypeOf(model).Elem()
	mValue := reflect.ValueOf(model).Elem()

	for i := 0; i < mValue.NumField(); i++ {
		typeField := mType.Field(i)
		tag := typeField.Tag.Get("gforms")
		if tag == "" {
			tag = typeField.Name
		} else if tag == "-" {
			continue
		}
		v, ok := fi.CleanedData[tag]
		if ok {
			valueField := mValue.Field(i)
			switch valueField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, ok := v.(int)
				if !ok {
					value = 0
				}
				valueField.SetInt(int64(value))
			case reflect.Float32, reflect.Float64:
				value, ok := v.(float64)
				if !ok {
					value = 0.0
				}
				valueField.SetFloat(value)
			case reflect.String:
				value, ok := v.(string)
				if !ok {
					va, ok := v.([]string)
					if ok && len(va) > 0 {
						value = va[0]
					} else {
						value = ""
					}
				}
				valueField.SetString(value)
			case reflect.Slice:
				value, ok := v.([]string)
				if !ok {
					value = []string{}
				}
				valueField.Set(reflect.ValueOf(value))
			case reflect.Bool:
				value, ok := v.(bool)
				if !ok {
					value = false
				}
				valueField.SetBool(value)
			}
		}
	}
}
