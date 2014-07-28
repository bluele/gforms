package gforms

import (
	"bytes"
	"net/http"
	"reflect"
)

type FormFields []Field

type CleanedData map[string]interface{}

type FormInstance struct {
	Fields      FormFields
	Data        Data
	CleanedData CleanedData
	Errors      map[string]string
}

type Form func() *FormInstance

func DefineForm(fields FormFields) Form {
	return func() *FormInstance {
		f := FormInstance{
			Fields: fields,
		}
		return &f
	}
}

func (self *FormInstance) IsValid() bool {
	isValid := true
	cleanedData := CleanedData{}
	errors := map[string]string{}

	for _, field := range self.Fields {
		name := field.GetName()
		cleanedValue, err := field.Clean(self.Data)
		if err != nil {
			errors[name] = err.Error()
			isValid = false
			break
		}
		err = field.Validate(cleanedValue)
		if err != nil {
			errors[name] = err.Error()
			isValid = false
			break
		}
		refV := reflect.ValueOf(cleanedValue)
		if !refV.IsValid() {
			cleanedData[name] = nil
			continue
		}
		el := refV.Elem()
		var v interface{}
		switch el.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = el.Int()
		case reflect.Float32, reflect.Float64:
			v = el.Float()
		case reflect.String:
			v = el.String()
		}
		cleanedData[name] = v
	}

	if isValid {
		self.CleanedData = cleanedData
	} else {
		self.Errors = errors
	}
	return isValid
}

func (self *FormInstance) ParseRequest(req *http.Request) error {
	req.ParseForm()
	data, err := parseReuqestBody(req)
	if err != nil {
		return err
	}
	self.Data = *data
	return nil
}

func (self *FormInstance) Html() string {
	var html bytes.Buffer
	for _, field := range self.Fields {
		html.WriteString(field.Html() + "\n")
	}
	return html.String()
}

func (self *FormInstance) MapTo(model interface{}) {
	if self.CleanedData == nil {
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
		v, ok := self.CleanedData[tag]
		if ok {
			valueField := mValue.Field(i)
			switch valueField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, ok := v.(int64)
				if !ok {
					value = 0
				}
				valueField.SetInt(value)
			case reflect.Float32, reflect.Float64:
				value, ok := v.(float64)
				if !ok {
					value = 0.0
				}
				valueField.SetFloat(value)
			case reflect.String:
				value, ok := v.(string)
				if !ok {
					value = ""
				}
				valueField.SetString(value)
			}
		}
	}
}
