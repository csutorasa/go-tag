package gotagio

import (
	"reflect"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that applies another [gotag.ValueWriterFunc] for all values in a slice and sets the value to the [reflect.Value].
// S generic type is the item type in the source slice.
func NewMapValueWriter[K comparable, V any](keyValueWriter gotag.ValueWriterFunc[K], valueValueWriter gotag.ValueWriterFunc[V]) gotag.ValueWriterFunc[map[K]V] {
	return func(v reflect.Value, source map[K]V) (bool, error) {
		if !v.CanSet() || v.Kind() != reflect.Map {
			return false, nil
		}
		s := reflect.MakeMapWithSize(v.Type(), len(source))
		for key, value := range source {
			keyValue := reflect.New(v.Type().Key()).Elem()
			supports, err := keyValueWriter(keyValue, key)
			if !supports {
				return false, nil
			}
			if err != nil {
				return true, gotag.NewWriteValueError(key, keyValue.Type(), err)
			}
			valueValue := reflect.New(v.Type().Elem()).Elem()
			supports, err = valueValueWriter(valueValue, value)
			if !supports {
				return false, nil
			}
			if err != nil {
				return true, gotag.NewWriteValueError(value, valueValue.Type(), err)
			}
			s.SetMapIndex(keyValue, valueValue)
		}
		v.Set(s)
		return true, nil
	}
}

// [gotag.ValueReaderFunc] that applies another [gotag.ValueReaderFunc] for keys and values in a map and gets them from the [reflect.Value].
// K generic type is the key type in the result map, V generic type is the value type in the result map.
func NewMapValueReader[K comparable, V any](keyValueReader gotag.ValueReaderFunc[K], valueValueReader gotag.ValueReaderFunc[V]) gotag.ValueReaderFunc[map[K]V] {
	return func(v reflect.Value) (map[K]V, bool, error) {
		if v.Kind() != reflect.Map {
			return nil, false, nil
		}
		l := v.Len()
		s := make(map[K]V, l)
		for _, keyValue := range v.MapKeys() {
			key, supports, err := keyValueReader(keyValue)
			if !supports {
				return nil, false, nil
			}
			if err != nil {
				return nil, true, gotag.NewReadValueError[K](keyValue.Type(), err)
			}
			valueValue := v.MapIndex(keyValue)
			value, supports, err := valueValueReader(valueValue)
			if !supports {
				return nil, false, nil
			}
			if err != nil {
				return nil, true, gotag.NewReadValueError[V](valueValue.Type(), err)
			}
			s[key] = value
		}
		return s, true, nil
	}
}
