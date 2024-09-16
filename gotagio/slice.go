package gotagio

import (
	"reflect"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that applies another [gotag.ValueWriterFunc] for all values in a slice and sets the value to the [reflect.Value].
// S generic type is the item type in the source slice.
func NewSliceValueWriter[S any](valueWriter gotag.ValueWriterFunc[S]) gotag.ValueWriterFunc[[]S] {
	return func(v reflect.Value, source []S) (bool, error) {
		if !v.CanSet() || v.Kind() != reflect.Slice {
			return false, nil
		}
		s := reflect.MakeSlice(v.Type(), len(source), len(source))
		for i := 0; i < len(source); i++ {
			sourceElement := source[i]
			sliceValue := s.Index(i)
			supports, err := valueWriter(sliceValue, sourceElement)
			if !supports {
				return false, nil
			}
			if err != nil {
				return true, gotag.NewWriteValueError(sourceElement, sliceValue.Type(), err)
			}
		}
		v.Set(s)
		return true, nil
	}
}

// [gotag.ValueReaderFunc] that applies another [gotag.ValueReaderFunc] for all values in a slice or array and gets them from the [reflect.Value].
// S generic type is the item type in the result slice.
func NewSliceValueReader[S any](valueReader gotag.ValueReaderFunc[S]) gotag.ValueReaderFunc[[]S] {
	return func(v reflect.Value) ([]S, bool, error) {
		if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
			return nil, false, nil
		}
		l := v.Len()
		s := make([]S, l)
		for i := 0; i < l; i++ {
			sliceValue := v.Index(i)
			result, supports, err := valueReader(sliceValue)
			if !supports {
				return nil, false, nil
			}
			if err != nil {
				return nil, true, gotag.NewReadValueError[[]S](sliceValue.Type(), err)
			}
			s[i] = result
		}
		return s, true, nil
	}
}
