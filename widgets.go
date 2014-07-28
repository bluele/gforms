package gforms

type Widget interface {
	Html(field Field) string
	Validate(value interface{}) error
}
