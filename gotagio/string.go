package gotagio

import (
	"fmt"
	"reflect"
)

// [gotag.ValueWriterFunc] that sets a string value to string [reflect.Value].
func WriteString(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || v.Kind() != reflect.String {
		return false, nil
	}
	v.SetString(s)
	return true, nil
}

// [gotag.ValueReaderFunc] that gets a string value from string [reflect.Value].
func ReadString(v reflect.Value) (string, bool, error) {
	if v.Kind() != reflect.String {
		return "", false, nil
	}
	return v.String(), true, nil
}

// [gotag.ValueWriterFunc] that sets the [fmt.Stringer] formatted string value to string [reflect.Value].
func WriteStringer(v reflect.Value, s fmt.Stringer) (bool, error) {
	return WriteString(v, s.String())
}
