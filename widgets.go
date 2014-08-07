package gforms

type Widget interface {
	html(Field, ...string) string
}
