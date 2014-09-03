package gforms

import (
	"bytes"
)

type checkboxMultipleWidget struct {
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
	Field   FieldInterface
	Attrs   map[string]string
	Options checkboxOptionValues
}

type CheckboxOptionsMaker func() CheckboxOptions

type CheckboxOptions interface {
	Label(int) string
	Value(int) string
	Checked(int) bool
	Disabled(int) bool
	Len() int
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

func (wg *checkboxMultipleWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	ctx := new(CheckboxContext)
	opts := wg.Maker()
	for i := 0; i < opts.Len(); i++ {
		ctx.Options = append(
			ctx.Options,
			&checkboxOptionValue{
				Label:    opts.Label(i),
				Value:    opts.Value(i),
				Checked:  opts.Checked(i),
				Disabled: opts.Disabled(i),
			})
	}
	ctx.Field = f
	ctx.Attrs = wg.Attrs
	err := Template.ExecuteTemplate(&buffer, "CheckboxMultipleWidget", ctx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate checkbox input field: <input type="checkbox" ...>
func CheckboxMultipleWidget(attrs map[string]string, mk CheckboxOptionsMaker) *checkboxMultipleWidget {
	wg := new(checkboxMultipleWidget)
	if attrs == nil {
		attrs = map[string]string{}
	}
	if isNilValue(mk) {
		mk = func() CheckboxOptions {
			return StringCheckboxOptions([][]string{})
		}
	}
	wg.Maker = mk
	wg.Attrs = attrs
	return wg
}
