package gforms

import (
	"reflect"
)

type RawData map[string]string

type Data map[string]*V

type V struct {
	RawValues []string
	// not ptr-value
	Value interface{}
	IsNil bool
	// value's kind
	Kind reflect.Kind
}

func newV(values []string, kind reflect.Kind) *V {
	v := new(V)
	v.RawValues = values
	v.Kind = kind
	v.IsNil = true
	return v
}

func nilV() *V {
	v := new(V)
	v.IsNil = true
	return v
}
