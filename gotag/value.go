package gotag

import (
	"reflect"
)

// Function that transforms the source value, and writes it to the [reflect.Value].
// Returns true and nil error, if the execution has succeeded.
// Returns true and non-nil error, if the execution has failed.
// Returns false and nil error, if the value is not supported.
type ValueWriterFunc[S any] func(value reflect.Value, source S) (bool, error)

// Writes the value without exposing the [reflect.Value].
func (f ValueWriterFunc[S]) WriteValue(c *StructTagCommand, source S) error {
	supports, err := f(c.fieldValue, source)
	if !supports {
		return NewStructTagHandlerUnsupportedFieldError(c)
	}
	if err != nil {
		return NewStructTagHandlerExecutionError(c, err)
	}
	return nil
}

// Combines multiple [gotag.ValueWriterFunc] into a single one.
// Each function is executed in natural order.
// Returns the result from the first supported function.
// Return false and nil error, if no function supports the value.
func NewFirstSupportedValueWriter[S any](writers ...ValueWriterFunc[S]) ValueWriterFunc[S] {
	return func(value reflect.Value, source S) (bool, error) {
		for _, writer := range writers {
			supports, err := writer(value, source)
			if supports {
				return true, err
			}
		}
		return false, nil
	}
}

// Combines multiple [gotag.ValueWriterFunc] into a single one.
// Each function is executed in natural order.
// Returns the result from the first supported function which executes without errors.
// Return false and nil error, if no function supports the value or all supported functions fail.
func NewFirstSucceedValueWriter[S any](writers ...ValueWriterFunc[S]) ValueWriterFunc[S] {
	return func(value reflect.Value, source S) (bool, error) {
		for _, writer := range writers {
			supports, err := writer(value, source)
			if supports && err == nil {
				return true, nil
			}
		}
		return false, nil
	}
}

// Function that transforms the [reflect.Value].
// Returns value, true and nil error, if the execution has succeeded.
// Returns default value, true and non-nil error, if the execution has failed.
// Returns default value, false and nil error, if the value is not supported.
type ValueReaderFunc[R any] func(value reflect.Value) (R, bool, error)

// Reads the value without exposing the [reflect.Value].
func (f ValueReaderFunc[R]) ReadValue(c *StructTagCommand) (R, error) {
	result, supports, err := f(c.fieldValue)
	if !supports {
		return result, NewStructTagHandlerUnsupportedFieldError(c)
	}
	if err != nil {
		return result, NewStructTagHandlerExecutionError(c, err)
	}
	return result, nil
}

// Combines multiple [gotag.ValueReaderFunc] into a single one.
// Each function is executed in natural order.
// Returns the result from the first supported function.
// Return default value, false and nil error, if no function supports the value.
func NewFirstSupportedValueReader[R any](readers ...ValueReaderFunc[R]) ValueReaderFunc[R] {
	return func(value reflect.Value) (R, bool, error) {
		for _, reader := range readers {
			result, supports, err := reader(value)
			if supports {
				return result, true, err
			}

		}
		var r R
		return r, false, nil
	}
}

// Combines multiple [gotag.ValueReaderFunc] into a single one.
// Each function is executed in natural order.
// Returns the result from the first supported function.
// Return default value, false and nil error, if no function supports the value or all supported functions fail.
func NewFirstSucceedValueReader[R any](readers ...ValueReaderFunc[R]) ValueReaderFunc[R] {
	return func(value reflect.Value) (R, bool, error) {
		for _, reader := range readers {
			result, supports, err := reader(value)
			if supports && err == nil {
				return result, true, nil
			}

		}
		var r R
		return r, false, nil
	}
}
