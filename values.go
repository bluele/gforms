package gforms

import (
	"reflect"
)

type Data map[string]*V

type V struct {
	RawStr   string
	RawValue interface{}
	// not ptr-value
	Value interface{}
	IsNil bool
	// value's kind
	Kind reflect.Kind
}

func newV(str string, value interface{}, kind reflect.Kind) *V {
	v := new(V)
	v.Kind = kind
	v.IsNil = true
	v.RawStr = str
	v.RawValue = value
	return v
}

func nilV(str string) *V {
	v := new(V)
	v.IsNil = true
	v.RawStr = str
	return v
}

func (self *V) rawValueAsString() *string {
	v, ok := self.RawValue.([]string)
	if ok {
		return &v[0]
	} else {
		return nil
	}
}

func (self *V) rawValueAsBool() bool {
	v, ok := self.RawValue.(bool)
	if ok {
		return v
	} else {
		return false
	}
}

func (self *V) rawValueAsStringArray() []string {
	v, ok := self.RawValue.([]string)
	if ok {
		return v
	} else {
		return []string{}
	}
}
