package gotaghttp

import (
	"errors"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP query
const TagQuery string = "query"

type queryStructTagValueWriter struct {
	writer      gotag.ValueWriterFunc[string]
	sliceWriter gotag.ValueWriterFunc[[]string]
}

// Tag implements StructTagHandler
func (w *queryStructTagValueWriter) Tag() string {
	return TagQuery
}

// Write implements StructTagValueWriter[*http.Request]
func (w *queryStructTagValueWriter) Write(command *gotag.StructTagCommand, r *http.Request) error {
	if len(command.TagValues()) == 0 {
		return gotag.NewTagValueHandlerConfigError(command, "no tag value found")
	}
	if len(command.TagValues()) > 1 {
		return gotag.NewTagValueHandlerConfigError(command, "too many tag values found")
	}
	paramName := command.TagValues()[0]
	queryValue := r.URL.Query()[paramName]
	if len(queryValue) == 1 {
		err := w.writer.WriteValue(command, queryValue[0])
		if err == nil {
			return nil
		}
		notSupported := &gotag.StructTagHandlerUnsupportedFieldError{}
		if !errors.As(err, &notSupported) {
			return err
		}
	}
	return w.sliceWriter.WriteValue(command, queryValue)
}

// Creates a new handler for HTTP query params.
// It uses [net/http] Request.URL.Query().
func NewQueryWriter(writer gotag.ValueWriterFunc[string]) *queryStructTagValueWriter {
	return &queryStructTagValueWriter{
		writer:      writer,
		sliceWriter: gotag.NewFirstSupportedValueWriter(gotagio.NewSliceValueWriter(writer), gotagio.NewArrayValueWriter(writer)),
	}
}

// Default QueryWriter.
var QueryWriter gotag.StructTagValueWriter[*http.Request] = NewQueryWriter(gotag.NewFirstSupportedValueWriter(gotagio.WriteString, gotagio.WriteStrConv))
