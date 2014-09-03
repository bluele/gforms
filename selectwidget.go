package gforms

import (
	"bytes"
	"strings"
)

type selectWidget struct {
	Multiple bool
	Attrs    map[string]string
	Maker    SelectOptionsMaker
	Widget
}

type selectOptionValue struct {
	Label    string
	Value    string
	Selected bool
	Disabled bool
}

type selectOptionsValues []*selectOptionValue

type selectContext struct {
	Multiple bool
	Field    FieldInterface
	Attrs    map[string]string
	Options  selectOptionsValues
}

type SelectOptionsMaker func() SelectOptions

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
	if strings.ToLower(selected) == "true" {
		return true
	} else {
		return false
	}
}

func (opt StringSelectOptions) Disabled(i int) bool {
	disabled := opt[i][3]
	if strings.ToLower(disabled) == "true" {
		return true
	} else {
		return false
	}
}

func (cs StringSelectOptions) Len() int {
	return len(cs)
}

func (wg *selectWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	context := new(selectContext)
	context.Field = f
	context.Multiple = wg.Multiple
	opts := wg.Maker()
	for i := 0; i < opts.Len(); i++ {
		context.Options = append(context.Options, &selectOptionValue{Label: opts.Label(i), Value: opts.Value(i), Selected: opts.Selected(i), Disabled: opts.Disabled(i)})
	}
	context.Attrs = wg.Attrs
	err := Template.ExecuteTemplate(&buffer, "SelectWidget", context)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate select and options field: <select><option></option></select>
func SelectWidget(attrs map[string]string, mk SelectOptionsMaker) *selectWidget {
	wg := new(selectWidget)
	if attrs == nil {
		attrs = map[string]string{}
	}
	if isNilValue(mk) {
		mk = func() SelectOptions {
			return StringSelectOptions([][]string{})
		}
	}
	wg.Maker = mk
	wg.Attrs = attrs
	return wg
}

// Generate select-multiple and options field: <select multiple><option></option></select>
func SelectMultipleWidget(attrs map[string]string, mk SelectOptionsMaker) *selectWidget {
	wg := SelectWidget(attrs, mk)
	wg.Multiple = true
	return wg
}
