package gforms

import (
	"bytes"
)

type selectOptionValue struct {
	Label    string
	Value    string
	Selected bool
	Disabled bool
}

type selectOptionsValues []*selectOptionValue

type SelectContext struct {
	Name    string
	Attrs   map[string]string
	Options selectOptionsValues
}

type SelectWidget struct {
	Attrs map[string]string
	Maker ChoiceMaker
	Widget
}

func (self *SelectWidget) html(field Field, vs ...string) string {
	var buffer bytes.Buffer
	context := new(SelectContext)
	opts := self.Maker()
	for i := 0; i < opts.Len(); i++ {
		context.Options = append(context.Options, &selectOptionValue{Label: opts.Label(i), Value: opts.Value(i), Selected: opts.Selected(i), Disabled: opts.Disabled(i)})
	}
	context.Name = field.GetName()
	context.Attrs = self.Attrs
	err := Template.ExecuteTemplate(&buffer, "SelectWidget", context)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type ChoiceMaker func() SelectOptions

type SelectOptions interface {
	Label(int) string
	Value(int) string
	Selected(int) bool
	Disabled(int) bool
	Len() int
}

type StringSelectOptions [][]string

func (opt StringSelectOptions) Label(i int) string {
	return opt[i][0]
}

func (opt StringSelectOptions) Value(i int) string {
	return opt[i][1]
}

func (opt StringSelectOptions) Selected(i int) bool {
	selected := opt[i][2]
	if selected == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringSelectOptions) Disabled(i int) bool {
	disabled := opt[i][3]
	if disabled == "true" {
		return true
	} else {
		return false
	}
}

func (cs StringSelectOptions) Len() int {
	return len(cs)
}

func NewSelectWidget(attrs map[string]string, cb ChoiceMaker) *SelectWidget {
	self := new(SelectWidget)
	self.Attrs = attrs
	self.Maker = cb
	return self
}
