package gotaghttp

import (
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP path params
const TagPathParam string = "pathParam"

type pathParamStructTagValueWriter struct {
	writer gotag.ValueWriterFunc[string]
}

// Tag implements StructTagHandler
func (w *pathParamStructTagValueWriter) Tag() string {
	return TagPathParam
}

// Write implements StructTagValueWriter[*http.Request]
func (w *pathParamStructTagValueWriter) Write(cmd *gotag.StructTagCommand, r *http.Request) error {
	if len(cmd.TagValues()) == 0 {
		return gotag.NewTagValueHandlerConfigError(cmd, "no tag value found")
	}
	if len(cmd.TagValues()) > 1 {
		return gotag.NewTagValueHandlerConfigError(cmd, "too many tag values found")
	}
	paramName := cmd.TagValues()[0]
	pathValue := r.PathValue(paramName)
	return w.writer.WriteValue(cmd, pathValue)
}

// Creates a new handler for HTTP path params.
// It uses [net/http] Request.PathValue.
func NewPathParamWriter(writer gotag.ValueWriterFunc[string]) *pathParamStructTagValueWriter {
	return &pathParamStructTagValueWriter{
		writer: writer,
	}
}

// Default PathParamWriter.
var PathParamWriter gotag.StructTagValueWriter[*http.Request] = NewPathParamWriter(gotag.NewValueWriters(gotagio.WriteString, gotagio.WriteStrConv))
