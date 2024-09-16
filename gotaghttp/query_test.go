package gotaghttp_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type QueryStringParams struct {
	Param1 string `queryParam:"p1"`
	Param2 string `queryParam:"p2"`
}

func TestQueryString(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryStringParams](gotaghttp.QueryWriter)
	err := doRequest("/", "/?p1=test&p2=value", "plain/text", []byte{}, testCreator, func(params QueryStringParams) {
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

type QueryNumberParams struct {
	Param1 float32 `queryParam:"p1"`
	Param2 int     `queryParam:"p2"`
	Param3 uint8   `queryParam:"p2"`
}

func TestQueryNumberString(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryNumberParams](gotaghttp.QueryWriter)
	err := doRequest("/", "/?p1=23%2E45&p2=234", "plain/text", []byte{}, testCreator, func(params QueryNumberParams) {
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

type QueryStringSliceParams struct {
	Param1 []string `queryParam:"p1"`
	Param2 []string `queryParam:"p2"`
}

func TestQueryStringSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryStringSliceParams](gotaghttp.QueryWriter)
	err := doRequest("/", "/?p1=test&p1=value", "plain/text", []byte{}, testCreator, func(params QueryStringSliceParams) {
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

type QueryNumberSliceParams struct {
	Param1 []float32 `queryParam:"p1"`
	Param2 []int     `queryParam:"p2"`
	Param3 []uint8   `queryParam:"p2"`
}

func TestQueryNumberSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[QueryNumberSliceParams](gotaghttp.QueryWriter)
	err := doRequest("/", "/?p1=23%2E45&p2=234", "plain/text", []byte{}, testCreator, func(params QueryNumberSliceParams) {
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
