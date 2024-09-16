package gotaghttp_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type QueryParamStringParams struct {
	Param1 string `queryParam:"p1"`
	Param2 string `queryParam:"p2"`
}

func TestQueryParamString(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryParamStringParams](gotaghttp.QueryParamWriter)
	err := doRequest("/", "/?p1=test&p2=value", "plain/text", []byte{}, testCreator, func(params QueryParamStringParams) {
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

type QueryParamNumberParams struct {
	Param1 float32 `queryParam:"p1"`
	Param2 int     `queryParam:"p2"`
	Param3 uint8   `queryParam:"p2"`
}

func TestQueryNumberString(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryParamNumberParams](gotaghttp.QueryParamWriter)
	err := doRequest("/", "/?p1=23%2E45&p2=234", "plain/text", []byte{}, testCreator, func(params QueryParamNumberParams) {
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

type QueryParamStringSliceParams struct {
	Param1 []string `queryParam:"p1"`
	Param2 []string `queryParam:"p2"`
}

func TestQueryParamStringSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryParamStringSliceParams](gotaghttp.QueryParamWriter)
	err := doRequest("/", "/?p1=test&p1=value", "plain/text", []byte{}, testCreator, func(params QueryParamStringSliceParams) {
		if len(params.Param1) != 2 {
			t.Fail()
		}
		if params.Param1[0] != "test" {
			t.Fail()
		}
		if params.Param1[1] != "value" {
			t.Fail()
		}
		if len(params.Param2) != 0 {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type QueryParamNumberSliceParams struct {
	Param1 []float32 `queryParam:"p1"`
	Param2 []int     `queryParam:"p2"`
	Param3 []uint8   `queryParam:"p2"`
}

func TestQueryParamNumberSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryParamNumberSliceParams](gotaghttp.QueryParamWriter)
	err := doRequest("/", "/?p1=23%2E45&p2=234", "plain/text", []byte{}, testCreator, func(params QueryParamNumberSliceParams) {
		if params.Param1[0] != 23.45 {
			t.Fail()
		}
		if params.Param2[0] != 234 {
			t.Fail()
		}
		if params.Param3[0] != 234 {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}
