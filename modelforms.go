package gforms

import (
	"fmt"
	"net/http"
	"reflect"
)

type FieldContext struct {
	Name  string
	Type  reflect.Type
	Value reflect.Value
}

type ModelContext struct {
	ModelType     reflect.Type
	FieldContexts []FieldContext
}

type ModelFormInstance struct {
	FormInstance
	ModelContext
}

type ModelForm func(...*http.Request) *ModelFormInstance

func unknownPanic(c FieldContext) {
	panic(fmt.Sprintf("Error: Unknown field type: %v", c.Value.Kind()))
}

// Auto Generate fileds by struct model.
func (ctx ModelContext) generateFields() []Field {
	fields := []Field{}
	for _, c := range ctx.FieldContexts {
		var field Field
		switch c.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field = NewIntegerField(c.Name, nil)
		case reflect.Float32, reflect.Float64:
			field = NewFloatField(c.Name, nil)
		case reflect.String:
			field = NewTextField(c.Name, nil)
		case reflect.Bool:
			field = NewBooleanField(c.Name, nil)
		case reflect.Slice:
			switch c.Value.Type().Elem().Kind() {
			case reflect.String:
				field = NewMultipleTextField(c.Name, nil, nil)
			}
		case reflect.Struct:
			switch c.Type.String() {
			case "time.Time":
				field = NewDateTimeField(c.Name, DefaultDateTimeFormat, nil)
			}
		}
		if field != nil {
			fields = append(fields, field)
		}
	}
	return fields
}

func getModelContext(model interface{}) ModelContext {
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
			typeField.Type,
			mValue.Field(i),
		})
	}
	return c
}

func newModelFormInstance(fs *Fields, ctx ModelContext) ModelFormInstance {
	mfi := ModelFormInstance{
		ModelContext: ctx,
	}
	mfi.fieldInstances = newFieldInterfaces(fs)
	return mfi
}

// Define a new form with generating fields from model's attributes and specified fields.
func DefineModelForm(model interface{}, fs *Fields) ModelForm {
	ctx := getModelContext(model)
	if fs == nil {
		fs = NewFields()
	}
	for _, f := range ctx.generateFields() {
		fs.AddField(f)
	}
	return func(r ...*http.Request) *ModelFormInstance {
		f := newModelFormInstance(fs, ctx)
		if len(r) > 0 {
			f.ParseError = f.parseRequest(r[0])
		}
		return &f
	}
}

func (mfi *ModelFormInstance) GetModel() interface{} {
	model := reflect.New(mfi.ModelType)
	mfi.MapTo(model.Interface())
	return model.Elem().Interface()
}
