package gotagio

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/csutorasa/go-tags/gotag"
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

// [gotag.ValueWriterFunc] that checks for a given set of true and false values and sets the bool value to the [reflect.Value].
// It panics if the true and false values have a common value.
func NewBoolWriter(trueValues []string, falseValues []string) gotag.ValueWriterFunc[string] {
	for _, v := range trueValues {
		if slices.Contains(falseValues, v) {
			panic("value is in both trueValues and falseValues")
		}
	}
	return func(v reflect.Value, s string) (bool, error) {
		if !v.CanSet() || v.Kind() != reflect.Bool {
			return false, nil
		}
		if slices.Contains(trueValues, s) {
			v.SetBool(true)
			return true, nil
		}
		if slices.Contains(falseValues, s) {
			v.SetBool(false)
			return true, nil
		}
		return false, gotag.NewWriteValueError(s, v.Type(), fmt.Errorf("%s is not in true or false values", s))
	}
}

// [gotag.ValueReaderFunc] that gets a string value from string [reflect.Value].
func ReadBool(v reflect.Value) (string, bool, error) {
	if v.Kind() != reflect.String {
		return "", false, nil
	}
	return v.String(), true, nil
}

// [gotag.ValueWriterFunc] that sets the [fmt.Stringer] formatted string value to string [reflect.Value].
func WriteStringer(v reflect.Value, s fmt.Stringer) (bool, error) {
	return WriteString(v, s.String())
}
