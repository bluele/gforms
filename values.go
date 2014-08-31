package gforms

import (
	"mime/multipart"
	"reflect"
)

type RawData map[string]string

type Data map[string]*V

type V struct {
	RawValue interface{}
	// not ptr-value
	Value interface{}
	IsNil bool
	// value's kind
	Kind reflect.Kind
}

func (self *V) rawValueAsString() *string {
	v, ok := self.RawValue.([]string)
	if ok {
		return &v[0]
	} else {
		return nil
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

func (self *V) rawValueAsFileHeader() multipart.FileHeader {
	headers, ok := self.RawValue.([]*multipart.FileHeader)
	if ok {
		return *headers[0]
	}
	return multipart.FileHeader{}
}

func (self *V) rawValueAsBool() bool {
	v, ok := self.RawValue.(bool)
	if ok {
		return v
	} else {
		return false
	}
}

func newV(value interface{}, kind reflect.Kind) *V {
	v := new(V)
	v.RawValue = value
	v.Kind = kind
	v.IsNil = true
	return v
}

func nilV() *V {
	v := new(V)
	v.IsNil = true
	return v
}
