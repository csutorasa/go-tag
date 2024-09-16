package gotagio

import (
	"reflect"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that applies another [gotag.ValueWriterFunc] for all values in an array and sets the value to the [reflect.Value].
// S generic type is the item type in the source slice.
func NewArrayValueWriter[S any](writer gotag.ValueWriterFunc[S]) gotag.ValueWriterFunc[[]S] {
	return func(v reflect.Value, source []S) (bool, error) {
		if v.Kind() != reflect.Array {
			return false, nil
		}
		l := v.Type().Len()
		for i := 0; i < l && i < len(source); i++ {
			sourceElement := source[i]
			arrayValue := v.Index(i)
			supports, err := writer(arrayValue, sourceElement)
			if !supports {
				return false, nil
			}
			if err != nil {
				return true, gotag.NewWriteValueError(sourceElement, arrayValue.Type(), err)
			}
		}
		return true, nil
	}
}
