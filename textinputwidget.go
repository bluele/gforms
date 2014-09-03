package gforms

import (
	"bytes"
)

type textInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *textInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "SimpleWidget", widgetContext{
		Type:  wg.Type,
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate text input fiele: <input type="text" ...>
func TextInputWidget(attrs map[string]string) Widget {
	w := new(textInputWidget)
	w.Type = "text"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
