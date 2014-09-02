package gforms

type Widget interface {
	html(FieldInterface) string
}

type widgetContext struct {
	Type  string
	Field FieldInterface
	Value string
	Attrs map[string]string
}
