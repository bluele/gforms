package gforms

type Widget interface {
	Html(Field) string
}
