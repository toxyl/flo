package utils

import "reflect"

func IsPointer(i any) bool {
	return reflect.TypeOf(i).Kind() == reflect.Ptr
}

func Dereference(value interface{}) interface{} {
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		return rv.Elem().Interface()
	}
	return value
}
