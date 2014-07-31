package gforms

import (
	"bytes"
)

type RadioWidget struct {
	Attrs map[string]string
	Maker RadioOptionsMaker
	Widget
}

type radioOptionValue struct {
	Label    string
	Value    string
	Checked  bool
	Disabled bool
}

type radioOptionValues []*radioOptionValue

type RadioContext struct {
	Name    string
	Attrs   map[string]string
	Options radioOptionValues
}

func (self *RadioWidget) Html(field Field) string {
	var buffer bytes.Buffer
	cx := new(RadioContext)
	opts := self.Maker()
	for i := 0; i < opts.Len(); i++ {
		cx.Options = append(
			cx.Options,
			&radioOptionValue{
				Label:    opts.Label(i),
				Value:    opts.Value(i),
				Checked:  opts.Checked(i),
				Disabled: opts.Disabled(i),
			})
	}
	cx.Name = field.GetName()
	cx.Attrs = self.Attrs
	err := Template.ExecuteTemplate(&buffer, "RadioWidget", cx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type RadioOptionsMaker func() RadioOptions

type RadioOptions interface {
	Label(int) string
	Value(int) string
	Checked(int) bool
	Disabled(int) bool
	Len() int
}

func NewRadioWidget(attrs map[string]string, cb RadioOptionsMaker) *RadioWidget {
	self := new(RadioWidget)
	self.Attrs = attrs
	self.Maker = cb
	return self
}

type StringRadioOptions [][]string

func (opt StringRadioOptions) Label(i int) string {
	return opt[i][0]
}

func (opt StringRadioOptions) Value(i int) string {
	return opt[i][1]
}

func (opt StringRadioOptions) Checked(i int) bool {
	checked := opt[i][2]
	if checked == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringRadioOptions) Disabled(i int) bool {
	disabled := opt[i][3]
	if disabled == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringRadioOptions) Len() int {
	return len(opt)
}
