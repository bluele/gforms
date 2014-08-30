package gforms

import (
	"bytes"
)

type PasswordWigdet struct {
	Attrs map[string]string
	Widget
}

func (self *PasswordWigdet) html(field Field, vs ...string) string {
	var buffer bytes.Buffer
	var v string
	if len(vs) > 0 {
		v = vs[0]
	}
	Template.ExecuteTemplate(&buffer, "SimpleWidget", widgetContext{
		Type:  "password",
		Attrs: self.Attrs,
		Name:  field.GetName(),
		Value: v,
	})
	return buffer.String()
}

func (self *PasswordWigdet) Validate(value interface{}) error {
	return nil
}

func NewPasswordWidget(attrs map[string]string) *PasswordWigdet {
	w := new(PasswordWigdet)
	w.Attrs = attrs
	return w
}
