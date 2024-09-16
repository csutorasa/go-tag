package gotagio

import (
	"fmt"
	"reflect"
	"time"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that parses time and set it to the [reflect.Value].
func NewTimeWriter(layout string) gotag.ValueWriterFunc[string] {
	return func(v reflect.Value, s string) (bool, error) {
		if !v.CanSet() || v.Type() != reflect.TypeFor[time.Time]() {
			return false, nil
		}
		t, err := time.Parse(layout, s)
		if err != nil {
			return true, gotag.NewWriteValueError(s, v.Type(), err)
		}
		v.Set(reflect.ValueOf(t))
		return true, nil
	}
}

// [gotag.ValueReaderFunc] that formats time and from the [reflect.Value].
func NewTimeReader(layout string) gotag.ValueReaderFunc[string] {
	return func(v reflect.Value) (string, bool, error) {
		if !v.CanInterface() || v.Type() != reflect.TypeFor[time.Time]() {
			return "", false, nil
		}
		t, ok := v.Interface().(time.Time)
		if !ok {
			return "", true, gotag.NewReadValueError[string](v.Type(), fmt.Errorf("invalid time value"))
		}
		return t.Format(layout), true, nil
	}
}

// [gotag.ValueWriterFunc] that parses duration and set it to the [reflect.Value].
func WriteDuration(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || v.Type() != reflect.TypeFor[time.Duration]() {
		return false, nil
	}
	t, err := time.ParseDuration(s)
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.Set(reflect.ValueOf(t))
	return true, nil
}

// [gotag.ValueReaderFunc] that formats duration and from the [reflect.Value].
func ReadDuration(v reflect.Value) (string, bool, error) {
	if !v.CanInterface() || v.Type() != reflect.TypeFor[time.Duration]() {
		return "", false, nil
	}
	d, ok := v.Interface().(time.Duration)
	if !ok {
		return "", true, gotag.NewReadValueError[string](v.Type(), fmt.Errorf("invalid duration value"))
	}
	return d.String(), true, nil
}
