package gforms

import (
	"bytes"
)

type TextField struct {
	BaseField
}

func (self *TextField) Html() string {
	if self.Widget == nil {
		return self.html()
	} else {
		return self.Widget.Html(self)
	}
}

func (self *TextField) html() string {
	var buffer bytes.Buffer
	Template.ExecuteTemplate(&buffer, "TextTypeField", self)
	return buffer.String()
}

func NewTextField(name string, vs Validators, ws ...Widget) *TextField {
	self := new(TextField)
	self.name = name
	self.validators = vs
	if len(ws) > 0 {
		self.Widget = ws[0]
	}
	return self
}
