package gforms

import (
    "bytes"
)

type emailInputWidget struct {
    Type  string
    Attrs map[string]string
    Widget
}

func (wg *emailInputWidget) html(f FieldInterface) string {
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

// Generate email input field: <input type="email" ...>
func EmailInputWidget(attrs map[string]string) Widget {
    w := new(emailInputWidget)
    w.Type = "email"
    if attrs == nil {
        attrs = map[string]string{}
    }
    w.Attrs = attrs
    return w
}
