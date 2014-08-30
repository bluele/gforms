package gforms

import (
	"reflect"
)

type FileField struct {
	BaseField
}

func (self *FileField) Html(rds ...RawData) string {
	return fieldToHtml(self, rds...)
}

func (self *FileField) html(vs ...string) string {
	return renderTemplate("FileTypeField", newTemplateContext(self, vs...))
}

// Create a new field for file object.
func NewFileField(name string, vs Validators, ws ...Widget) *FileField {
	self := new(FileField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}

func (self *FileField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := m.rawValueAsFileHeader()
		m.Kind = reflect.String
		m.Value = v
		m.IsNil = false
		return m, nil
	}
	return nilV(), nil
}
