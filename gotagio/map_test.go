package gotagio_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestMapValueWriter(t *testing.T) {
	var s map[int64]int64
	slice := gotagio.NewMapValueWriter(testInt64ValueWriter, testInt64ValueWriter)
	writeValue(t, slice, &s, map[string]string{"10": "20"})
	if len(s) != 1 {
		t.Fatal("length should remain the same")
	}
	v, ok := s[10]
	if !ok {
		t.Fatal("key should exist")
	}
	if v != 20 {
		t.Fatal("value should be the same")
	}
}

func TestMapValueReader(t *testing.T) {
	s := map[int64]int64{10: 20}
	slice := gotagio.NewMapValueReader(testInt64ValueReader, testInt64ValueReader)
	result := readValue(t, slice, s)
	if len(result) != 1 {
		t.Fatal("length should remain the same")
	}
	v, ok := s[10]
	if !ok {
		t.Fatal("key should exist")
	}
	if v != 20 {
		t.Fatal("value should be the same")
	}
}
