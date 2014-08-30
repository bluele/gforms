package gforms

type Widget interface {
	html(Field, ...string) string
}

type widgetContext struct {
	Type  string
	Name  string
	Value string
	Attrs map[string]string
}
