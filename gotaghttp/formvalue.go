package gotaghttp

import (
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP form values
const TagFormValue string = "formValue"

type formValueStructTagValueWriter struct {
	writer gotag.ValueWriterFunc[string]
}

// Tag implements StructTagHandler
func (w *formValueStructTagValueWriter) Tag() string {
	return TagFormValue
}

// Write implements StructTagValueWriter[*http.Request]
func (w *formValueStructTagValueWriter) Write(cmd *gotag.StructTagCommand, r *http.Request) error {
	if len(cmd.TagValues()) == 0 {
		return gotag.NewTagValueHandlerConfigError(cmd, "no tag value found")
	}
	if len(cmd.TagValues()) > 1 {
		return gotag.NewTagValueHandlerConfigError(cmd, "too many tag values found")
	}
	err := r.ParseForm()
	if err != nil {
		return gotag.NewStructTagHandlerExecutionError(cmd, err)
	}
	paramName := cmd.TagValues()[0]
	queryValue := r.PostFormValue(paramName)
	return w.writer.WriteValue(cmd, queryValue)
}

// Creates a new handler for HTTP form values.
// It uses [net/http] Request.ParseForm and Request.PostFormValue.
func NewFormValueWriter(writer gotag.ValueWriterFunc[string]) *formValueStructTagValueWriter {
	return &formValueStructTagValueWriter{
		writer: writer,
	}
}

// Default FormValueWriter.
var FormValueWriter gotag.StructTagValueWriter[*http.Request] = NewFormValueWriter(gotag.NewValueWriters(gotagio.WriteString, gotagio.WriteStrConv))
