package gforms

type Errors map[string][]string

func (es Errors) Has(key string) bool {
	_, ok := es[key]
	return ok
}

func (es Errors) Get(key string) []string {
	v, _ := es[key]
	return v
}
