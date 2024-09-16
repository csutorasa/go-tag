package gotaghttp_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type PathParamStringParams struct {
	Param1 string `pathParam:"p1"`
	Param2 string `pathParam:"p2"`
}

func TestPathParamString(t *testing.T) {
	testCreator := gotag.NewDecoder[PathParamStringParams](gotaghttp.PathParamWriter)
	err := doRequest("/{p1}/{p2}/23%2E45/234", "/test/value/23%2E45/234", "plain/text", []byte{}, testCreator, func(params PathParamStringParams) {
		if params.Param1 != "test" {
			t.Fail()
		}
		if params.Param2 != "value" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type PathParamNumberParams struct {
	Param1 float32 `pathParam:"p1"`
	Param2 int     `pathParam:"p2"`
	Param3 uint8   `pathParam:"p2"`
}

func TestPathNumberString(t *testing.T) {
	testCreator := gotag.NewDecoder[PathParamNumberParams](gotaghttp.PathParamWriter)
	err := doRequest("/test/value/{p1}/{p2}", "/test/value/23%2E45/234", "plain/text", []byte{}, testCreator, func(params PathParamNumberParams) {
		if params.Param1 != 23.45 {
			t.Fail()
		}
		if params.Param2 != 234 {
			t.Fail()
		}
		if params.Param3 != 234 {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}
