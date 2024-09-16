package gotaghttp

import (
	"net/http"

	"github.com/csutorasa/go-tags/gotag"
)

// Default HTTP writers.
var DefaultWriters = []gotag.StructTagValueWriter[*http.Request]{
	BodyWriter,
	FormValueWriter,
	PathParamWriter,
	QueryParamWriter,
}
