package gforms

import (
	"bytes"
	"reflect"
)

type CheckboxWidget struct {
	Attrs map[string]string
	Maker CheckboxOptionsMaker
	Widget
}

type checkboxOptionValue struct {
	Label    string
	Value    string
	Checked  bool
	Disabled bool
}

type checkboxOptionValues []*checkboxOptionValue

type CheckboxContext struct {
	Name    string
	Attrs   map[string]string
	Options checkboxOptionValues
}

func (self *CheckboxWidget) Html(field Field) string {
	var buffer bytes.Buffer
	cx := new(CheckboxContext)
	opts := self.Maker()
	for i := 0; i < opts.Len(); i++ {
		cx.Options = append(
			cx.Options,
			&checkboxOptionValue{
				Label:    opts.Label(i),
				Value:    opts.Value(i),
				Checked:  opts.Checked(i),
				Disabled: opts.Disabled(i),
			})
	}
	cx.Name = field.GetName()
	cx.Attrs = self.Attrs
	err := Template.ExecuteTemplate(&buffer, "CheckboxWidget", cx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (self *CheckboxWidget) Clean(name string, data Data) (*V, error) {
	m, hasField := data[name]
	if hasField {
		m.Kind = reflect.Slice
		m.Value = m.RawValues
		m.IsNill = false
		return m, nil
	}
	return nilV(), nil
}

type CheckboxOptionsMaker func() CheckboxOptions

type CheckboxOptions interface {
	Label(int) string
	Value(int) string
	Checked(int) bool
	Disabled(int) bool
	Len() int
}

func NewCheckboxWidget(attrs map[string]string, cb CheckboxOptionsMaker) *CheckboxWidget {
	self := new(CheckboxWidget)
	self.Attrs = attrs
	self.Maker = cb
	return self
}

type StringCheckboxOptions [][]string

func (opt StringCheckboxOptions) Label(i int) string {
	return opt[i][0]
}

func (opt StringCheckboxOptions) Value(i int) string {
	return opt[i][1]
}

func (opt StringCheckboxOptions) Checked(i int) bool {
	checked := opt[i][2]
	if checked == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringCheckboxOptions) Disabled(i int) bool {
	disabled := opt[i][3]
	if disabled == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringCheckboxOptions) Len() int {
	return len(opt)
}
