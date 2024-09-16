package gotagio_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/csutorasa/go-tags/gotag"
)

func writeValue[S any, V any](t *testing.T, c gotag.ValueWriterFunc[V], s *S, fieldValue V) {
	supported, err := c(reflect.ValueOf(s).Elem(), fieldValue)
	if !supported {
		t.Fatalf("ValueWriter should support fields")
	}
	if err != nil {
		t.Fatalf("ValueWriter failed %s", err.Error())
	}
}

func testInt64ValueWriter(v reflect.Value, s string) (bool, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return true, err
	}
	v.SetInt(i)
	return true, nil
}

func readValue[R any](t *testing.T, c gotag.ValueReaderFunc[R], s any) R {
	v := reflect.ValueOf(s)
	result, supported, err := c(v)
	if !supported {
		t.Fatalf("ValueReader should support fields")
	}
	if err != nil {
		t.Fatalf("ValueReader failed %s", err.Error())
	}
	return result
}

func testInt64ValueReader(v reflect.Value) (string, bool, error) {
	return strconv.FormatInt(v.Int(), 10), true, nil
}
