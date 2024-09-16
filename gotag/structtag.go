package gotag

import (
	"reflect"
)

// Interface for struct tag handlers.
type StructTagHandler interface {
	Tag() string
}

// Interface for struct tag based value writers.
type StructTagValueWriter[S any] interface {
	StructTagHandler
	// Processes the struct tag and sets the value.
	Write(cmd *StructTagCommand, source S) error
}

// Interface for struct tag based value writers.
type StructTagValueReader[R any] interface {
	StructTagHandler
	// Processes the struct tag and sets the value.
	Read(cmd *StructTagCommand, result R) (R, error)
}

// Allows accessing the value without exposing the [reflect.Value].
type StructTagCommand struct {
	tag         string
	tagValues   []string
	structType  reflect.Type
	structField reflect.StructField
	fieldValue  reflect.Value
}

// Tag implements StructTagHandler.
func (c *StructTagCommand) Tag() string {
	return c.tag
}

func (c *StructTagCommand) TagValues() []string {
	return c.tagValues
}
