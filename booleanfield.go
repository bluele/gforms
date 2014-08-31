package gforms

import (
	"bytes"
	"reflect"
)

type BooleanField struct {
	BaseField
}

type booleanContext struct {
	Name  string
	Value string
}

func (self *BooleanField) Html(rds ...RawData) string {
	return fieldToHtml(self, rds...)
}

func (self *BooleanField) html(vs ...string) string {
	var buffer bytes.Buffer
	cx := new(booleanContext)
	cx.Name = self.GetName()
	err := Template.ExecuteTemplate(&buffer, "BooleanTypeField", cx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (self *BooleanField) Clean(data Data) (*V, error) {
	m, hasField := data[self.GetName()]
	if hasField {
		v := false
		if m.Kind == reflect.String {
			vs := m.rawValueAsString()
			if vs != nil {
				v = true
			}
		} else if m.Kind == reflect.Bool {
			v = m.rawValueAsBool()
		}
		m.Value = v
		m.Kind = reflect.Bool
		m.IsNil = false
		return m, nil
	}
	nv := newV(false, reflect.Bool)
	nv.Value = false
	nv.IsNil = false
	return nv, nil
}

// Create a new field for boolean value.
func NewBooleanField(name string, vs Validators, ws ...Widget) *BooleanField {
	self := new(BooleanField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}
