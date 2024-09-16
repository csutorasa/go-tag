package gotag

import (
	"fmt"
	"reflect"
)

// An error that says that the type where the tags should be read is not a struct.
var ErrNotAStruct = fmt.Errorf("type is not a struct")

// An error that wraps the underlying value write error.
type WriteValueError struct {
	newValue      string
	valueTypeName string
	cause         error
}

// Creates a new [gotag.WriteValueError].
func NewWriteValueError(newValue any, valueType reflect.Type, cause error) *WriteValueError {
	return &WriteValueError{
		newValue:      fmt.Sprintf("%v", newValue),
		valueTypeName: typeName(valueType),
		cause:         cause,
	}
}

// Error implements error.
func (e *WriteValueError) Error() string {
	return fmt.Sprintf("failed to convert %v to %s: %s", e.newValue, e.valueTypeName, e.cause.Error())
}

// Returns the cause of the the write error.
func (e *WriteValueError) Unwrap() error {
	return e.cause
}

// An error that wraps the underlying value write error.
type ReadValueError struct {
	resultTypeName string
	valueTypeName  string
	cause          error
}

// Creates a new [gotag.ReadValueError].
func NewReadValueError[T any](valueType reflect.Type, cause error) *ReadValueError {
	return &ReadValueError{
		resultTypeName: typeName(reflect.TypeFor[T]()),
		valueTypeName:  typeName(valueType),
		cause:          cause,
	}
}

// Error implements error.
func (e *ReadValueError) Error() string {
	return fmt.Sprintf("failed to convert from %s to %s: %s", e.valueTypeName, e.resultTypeName, e.cause.Error())
}

// Returns the cause of the the read error.
func (e *ReadValueError) Unwrap() error {
	return e.cause
}

// An error that wraps the underlying value tag processing error.
type StructTagHandlerExecutionError struct {
	tag             string
	structTypeName  string
	structFieldName string
	cause           error
}

// Creates a new [gotag.StructTagHandlerExecutionError].
func NewStructTagHandlerExecutionError(command *StructTagCommand, cause error) *StructTagHandlerExecutionError {
	return &StructTagHandlerExecutionError{
		tag:             command.Tag(),
		structTypeName:  command.structType.Name(),
		structFieldName: command.structField.Name,
		cause:           cause,
	}
}

// Error implements error.
func (e *StructTagHandlerExecutionError) Error() string {
	return fmt.Sprintf("execution failed for field %s.%s and tag %s: %s", e.structTypeName, e.structFieldName, e.tag, e.cause.Error())
}

// Returns the cause of the the write error.
func (e *StructTagHandlerExecutionError) Unwrap() error {
	return e.cause
}

// An error for a not supported field.
type StructTagHandlerUnsupportedFieldError struct {
	tag             string
	structTypeName  string
	structFieldName string
}

// Creates a new [gotag.StructTagHandlerUnsupportedFieldError].
func NewStructTagHandlerUnsupportedFieldError(command *StructTagCommand) *StructTagHandlerUnsupportedFieldError {
	return &StructTagHandlerUnsupportedFieldError{
		tag:             command.Tag(),
		structTypeName:  command.structType.Name(),
		structFieldName: command.structField.Name,
	}
}

// Error implements error.
func (e *StructTagHandlerUnsupportedFieldError) Error() string {
	return fmt.Sprintf("unsupported field %s.%s for tag %s", e.structTypeName, e.structFieldName, e.tag)
}

// Creates a new [gotag.TagValueHandlerConfigError].
func NewTagValueHandlerConfigError(command *StructTagCommand, message string) *TagValueHandlerConfigError {
	return &TagValueHandlerConfigError{
		tag:             command.Tag(),
		message:         message,
		structTypeName:  command.structType.Name(),
		structFieldName: command.structField.Name,
	}
}

// An error for a invalid handler configuration.
type TagValueHandlerConfigError struct {
	tag             string
	structTypeName  string
	structFieldName string
	message         string
}

// Error implements error.
func (e *TagValueHandlerConfigError) Error() string {
	return fmt.Sprintf("invalid handler configuration for field %s.%s and tag %s: %s", e.structTypeName, e.structFieldName, e.tag, e.message)
}

func typeName(t reflect.Type) string {
	name := t.Name()
	if name == "" {
		return t.Kind().String()
	}
	return name
}
