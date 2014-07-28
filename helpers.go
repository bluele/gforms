package gforms

import (
	"reflect"
)

func isNilValue(value interface{}) bool {
	return ((value == nil) || reflect.ValueOf(value).IsNil())
}
