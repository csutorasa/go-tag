package gotaghttp

import (
	"errors"
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotagio"
)

// Struct tag for HTTP query params
const TagQueryParam string = "queryParam"

type queryParamStructTagValueWriter struct {
	writer      gotag.ValueWriterFunc[string]
	sliceWriter gotag.ValueWriterFunc[[]string]
}

// Tag implements StructTagHandler
func (w *queryParamStructTagValueWriter) Tag() string {
	return TagQueryParam
}

// Write implements StructTagValueWriter[*http.Request]
func (w *queryParamStructTagValueWriter) Write(command *gotag.StructTagCommand, r *http.Request) error {
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
func NewQueryParamWriter(writer gotag.ValueWriterFunc[string]) *queryParamStructTagValueWriter {
	return &queryParamStructTagValueWriter{
		writer:      writer,
		sliceWriter: gotag.NewValueWriters(gotagio.NewSliceValueWriter(writer), gotagio.NewArrayValueWriter(writer)),
	}
}

// Default QueryParamWriter.
var QueryParamWriter gotag.StructTagValueWriter[*http.Request] = NewQueryParamWriter(gotag.NewValueWriters(gotagio.WriteString, gotagio.WriteStrConv))
