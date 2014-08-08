package gforms

import (
	"bytes"
)

type selectOptionValue struct {
	Label string
	Value string
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
		context.Options = append(context.Options, &selectOptionValue{Label: opts.Label(i), Value: opts.Value(i)})
	}
	context.Name = field.GetName()
	context.Attrs = self.Attrs
	Template.ExecuteTemplate(&buffer, "SelectWidget", context)
	return buffer.String()
}

type ChoiceMaker func() SelectOptions

type SelectOptions interface {
	Label(int) string
	Value(int) string
	Len() int
}

type StringSelectOptions [][]string

func (cs StringSelectOptions) Label(i int) string {
	return cs[i][0]
}

func (cs StringSelectOptions) Value(i int) string {
	return cs[i][1]
}

func (cs StringSelectOptions) Len() int {
	return len(cs)
}

// [{"one", "1", "two", "2"}]
func NewSelectWidget(attrs map[string]string, cb ChoiceMaker) *SelectWidget {
	self := new(SelectWidget)
	self.Attrs = attrs
	self.Maker = cb
	return self
}
