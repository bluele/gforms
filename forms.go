package gforms

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
)

// Form fields list.
type Fields struct {
	fields    []Field
	fieldsMap map[string]Field
}

func (self *Fields) Index(i int) Field {
	return self.fields[i]
}

func (self *Fields) NamedBy(name string) (Field, bool) {
	v, ok := self.fieldsMap[name]
	return v, ok
}

func (self *Fields) GetMap() map[string]Field {
	return self.fieldsMap
}

func (self *Fields) AddField(field Field) bool {
	name := field.GetName()
	_, exists := self.NamedBy(name)
	if !exists {
		self.fields = append(self.fields, field)
		self.fieldsMap[name] = field
		return true
	}
	return false
}

func NewFields(fields ...Field) *Fields {
	fs := Fields{}
	fs.fieldsMap = make(map[string]Field)
	for _, field := range fields {
		fs.fieldsMap[field.GetName()] = field
	}
	fs.fields = fields
	return &fs
}

// cleaned data for all fields.
type CleanedData map[string]interface{}

type FormInstance struct {
	Fields      *Fields
	Data        Data
	RawData     RawData
	CleanedData CleanedData
	Errors      Errors
}

type FieldContext struct {
	Name  string
	Value reflect.Value
}

type ModelContext struct {
	ModelType     reflect.Type
	FieldContexts []FieldContext
}

// Auto Generate fileds by struct model.
func (self ModelContext) generateFields() []Field {
	fields := []Field{}
	for _, c := range self.FieldContexts {
		var field Field
		switch c.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field = NewIntegerField(c.Name, nil)
		case reflect.Float32, reflect.Float64:
			field = NewFloatField(c.Name, nil)
		case reflect.String:
			field = NewTextField(c.Name, nil)
		default:
			panic(fmt.Sprintf("Error: Unknown field type: %v", c.Value.Kind()))
		}
		fields = append(fields, field)
	}
	return fields
}

func newModelContext(model interface{}) ModelContext {
	mType := reflect.TypeOf(model)
	mValue := reflect.ValueOf(model)
	for {
		if mType.Kind() == reflect.Ptr {
			mType = mType.Elem()
			mValue = mValue.Elem()
		} else {
			break
		}
	}
	c := ModelContext{}
	c.ModelType = mType
	c.FieldContexts = []FieldContext{}

	for i := 0; i < mValue.NumField(); i++ {
		typeField := mType.Field(i)
		tag := typeField.Tag.Get("gforms")
		if tag == "" {
			tag = typeField.Name
		} else if tag == "-" {
			continue
		}
		c.FieldContexts = append(c.FieldContexts, FieldContext{
			tag,
			mValue.Field(i),
		})
	}
	return c
}

type ModelFormInstance struct {
	FormInstance
	ModelContext
}

func newModelFormInstance(model interface{}, fields *Fields) ModelFormInstance {
	c := newModelContext(model)
	for _, gf := range c.generateFields() {
		fields.AddField(gf)
	}
	inst := ModelFormInstance{
		ModelContext: c,
	}
	inst.Fields = fields
	return inst
}

type Form func(...*http.Request) *FormInstance

// Initialize with http request.
func (f Form) FromRequest(r *http.Request) *FormInstance {
	return f(r)
}

// Intialize with map object.
func (f Form) FromValues(v url.Values) *FormInstance {
	fi := f()
	fi.parseValues(v)
	return fi
}

type ModelForm func(...*http.Request) *ModelFormInstance

// Initialize with http request.
func (f ModelForm) FromRequest(r *http.Request) *ModelFormInstance {
	return f(r)
}

// Intialize with map object.
func (f ModelForm) FromValues(v url.Values) *ModelFormInstance {
	fi := f(nil)
	fi.parseValues(v)
	return fi
}

// Define new form with specified fields.
func DefineForm(fields *Fields) Form {
	return func(r ...*http.Request) *FormInstance {
		f := FormInstance{
			Fields: fields,
		}
		if len(r) > 0 {
			f.parseRequest(r[0])
		}
		return &f
	}
}

// Define new form with generating fields from model's attributes and specified fields.
func DefineModelForm(model interface{}, fields *Fields) ModelForm {
	return func(r ...*http.Request) *ModelFormInstance {
		f := newModelFormInstance(model, fields)
		if len(r) > 0 {
			f.parseRequest(r[0])
		}
		return &f
	}
}

type Cleaner interface {
	Clean(string, Data) (*V, error)
}

// Check validation for forminstance data.
func (self *FormInstance) IsValid() bool {
	isValid := true
	cleanedData := CleanedData{}
	errors := map[string]string{}

	for _, field := range self.Fields.fields {
		name := field.GetName()
		widget := field.GetWigdet()
		var err error
		var cleanedValue *V
		if widget == nil {
			cleanedValue, err = field.Clean(self.Data)
		} else {
			if cleaner, ok := widget.(Cleaner); ok {
				cleanedValue, err = cleaner.Clean(name, self.Data)
			} else {
				cleanedValue, err = field.Clean(self.Data)
			}
		}

		if err != nil {
			errors[name] = err.Error()
			isValid = false
			continue
		}

		err = field.Validate(cleanedValue, cleanedData)
		if err != nil {
			errors[name] = err.Error()
			isValid = false
			continue
		}
		if !cleanedValue.IsNil {
			cleanedData[name] = cleanedValue.Value
		}
	}

	if isValid {
		self.CleanedData = cleanedData
	} else {
		self.Errors = errors
	}
	return isValid
}

func (self *FormInstance) parseRequest(req *http.Request) error {
	data, rawData, err := parseReuqestBody(req)
	if err != nil {
		return err
	}
	if data == nil || rawData == nil {
		return nil
	}
	self.Data = *data
	self.RawData = *rawData
	return nil
}

func (self *FormInstance) parseValues(v url.Values) error {
	data, rawData, err := bindValues(v)
	if err != nil {
		return err
	}
	if data == nil || rawData == nil {
		return nil
	}
	self.Data = *data
	self.RawData = *rawData
	return nil
}

func (self *FormInstance) Html() string {
	var html bytes.Buffer
	for _, field := range self.Fields.fields {
		html.WriteString(field.Html(self.RawData) + "\n")
	}
	return html.String()
}

// maps cleanedData to specified model.
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
			case reflect.Struct: // file
				value, ok := v.(multipart.FileHeader)
				if !ok {
					value = multipart.FileHeader{}
				}
				valueField.Set(reflect.ValueOf(value))
			}
		}
	}
}

func (self *ModelFormInstance) GetModel() interface{} {
	model := reflect.New(self.ModelType)
	self.MapTo(model.Interface())
	return model.Elem().Interface()
}
