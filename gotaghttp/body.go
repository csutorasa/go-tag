package gotaghttp

import (
	"fmt"
	"io"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP bodies
const TagBody string = "body"

type bodyStructTagValueWriter struct {
	converters map[string]gotag.ValueWriterFunc[io.Reader]
	writer     gotag.ValueWriterFunc[io.Reader]
}

// Tag implements StructTagHandler
func (w *bodyStructTagValueWriter) Tag() string {
	return TagBody
}

// Write implements StructTagValueWriter[*http.Request]
func (w *bodyStructTagValueWriter) Write(command *gotag.StructTagCommand, r *http.Request) error {
	if r.Body == nil {
		return nil
	}
	if len(command.TagValues()) > 1 {
		return gotag.NewTagValueHandlerConfigError(command, "too many tag values found")
	}
	if len(command.TagValues()) == 1 {
		converterValue := command.TagValues()[0]
		converter, ok := w.converters[converterValue]
		if !ok {
			return gotag.NewTagValueHandlerConfigError(command, fmt.Sprintf("unknown converter %s", converterValue))
		}
		return converter.WriteValue(command, io.Reader(r.Body))
	}
	return w.writer.WriteValue(command, io.Reader(r.Body))
}

// Creates a new handler for HTTP bodies.
// It uses [net/http] Request.Body.
func NewBodyWriter(converters map[string]gotag.ValueWriterFunc[io.Reader], writer gotag.ValueWriterFunc[io.Reader]) *bodyStructTagValueWriter {
	return &bodyStructTagValueWriter{
		converters: converters,
		writer:     writer,
	}
}

// Default BodyWriter.
var BodyWriter gotag.StructTagValueWriter[*http.Request] = NewBodyWriter(map[string]gotag.ValueWriterFunc[io.Reader]{
	"json": gotagio.WriteJsonReader,
	"xml":  gotagio.WriteXmlReader,
}, gotagio.WriteFromReader)
