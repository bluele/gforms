package gforms

import (
    "bytes"
)

type datetimeInputWidget struct {
    Type  string
    Attrs map[string]string
    Widget
}

func (wg *datetimeInputWidget) html(f FieldInterface) string {
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

// Generate text input field: <input type="text" ...>
func DatetimeInputWidget(attrs map[string]string) Widget {
    w := new(datetimeInputWidget)
    w.Type = "text"
    if attrs == nil {
        attrs = map[string]string{}
    }
    w.Attrs = attrs
    return w
}
