package gotaghttp_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type PathValueStringParams struct {
	Param1 string `pathValue:"p1"`
	Param2 string `pathValue:"p2"`
}

func TestPathValueString(t *testing.T) {
	testCreator := gotag.NewDecoder[PathValueStringParams](gotaghttp.PathValueWriter)
	err := doRequest("/{p1}/{p2}/23%2E45/234", "/test/value/23%2E45/234", "plain/text", []byte{}, testCreator, func(params PathValueStringParams) {
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

type PathValueNumberParams struct {
	Param1 float32 `pathValue:"p1"`
	Param2 int     `pathValue:"p2"`
	Param3 uint8   `pathValue:"p2"`
}

func TestPathValueNumber(t *testing.T) {
	testCreator := gotag.NewDecoder[PathValueNumberParams](gotaghttp.PathValueWriter)
	err := doRequest("/test/value/{p1}/{p2}", "/test/value/23%2E45/234", "plain/text", []byte{}, testCreator, func(params PathValueNumberParams) {
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
