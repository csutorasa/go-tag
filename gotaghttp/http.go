// Defines [gotags.StructTagHandler]s for managing a [net/http] Request.
package gotaghttp

import (
	"errors"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
)

// Default HTTP writers.
var DefaultWriters = []gotag.StructTagValueWriter[*http.Request]{
	BodyWriter,
	FormValueWriter,
	PathValueWriter,
	QueryWriter,
}

// Checks if the execution is due the invalid or unexpected data.
func IsExecutionError(err error) bool {
	executionErr := &gotag.StructTagHandlerExecutionError{}
	return errors.As(err, &executionErr)
}
