package gotaghttp

import (
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP path params
const TagPathValue string = "pathValue"

type pathValueStructTagValueWriter struct {
	writer gotag.ValueWriterFunc[string]
}

// Tag implements StructTagHandler
func (w *pathValueStructTagValueWriter) Tag() string {
	return TagPathValue
}

// Write implements StructTagValueWriter[*http.Request]
func (w *pathValueStructTagValueWriter) Write(cmd *gotag.StructTagCommand, r *http.Request) error {
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
func NewPathValueWriter(writer gotag.ValueWriterFunc[string]) *pathValueStructTagValueWriter {
	return &pathValueStructTagValueWriter{
		writer: writer,
	}
}

// Default PathValueWriter.
var PathValueWriter gotag.StructTagValueWriter[*http.Request] = NewPathValueWriter(gotag.NewFirstSupportedValueWriter(gotagio.WriteString, gotagio.WriteStrConv))
