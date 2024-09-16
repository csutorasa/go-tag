package gotaghttp_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type FormValueStringParams struct {
	Param1 string `formValue:"p1"`
	Param2 string `formValue:"p2"`
}

func TestFormValueString(t *testing.T) {
	testCreator := gotag.NewDecoder[FormValueStringParams](gotaghttp.FormValueWriter)
	err := doRequest("/", "/", "application/x-www-form-urlencoded", []byte("p1=test&p2=value"), testCreator, func(params FormValueStringParams) {
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

type FormValueNumberParams struct {
	Param1 float32 `formValue:"p1"`
	Param2 int     `formValue:"p2"`
	Param3 uint8   `formValue:"p2"`
}

func TestFormValueNumberString(t *testing.T) {
	testCreator := gotag.NewDecoder[FormValueNumberParams](gotaghttp.FormValueWriter)
	err := doRequest("/", "/", "application/x-www-form-urlencoded", []byte("p1=23.45&p2=234"), testCreator, func(params FormValueNumberParams) {
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

func TestFormValueIgnoresQueryParams(t *testing.T) {
	testCreator := gotag.NewDecoder[FormValueStringParams](gotaghttp.FormValueWriter)
	err := doRequest("/", "/?p1=test", "application/x-www-form-urlencoded", []byte("p2=value"), testCreator, func(params FormValueStringParams) {
		if params.Param1 != "" {
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
