package gotagio_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestSliceValueWriter(t *testing.T) {
	var s []int64
	slice := gotagio.NewSliceValueWriter(testInt64ValueWriter)
	writeValue(t, slice, &s, []string{"10", "20"})
	if len(s) != 2 {
		t.Fatal("length should remain the same")
	}
	if s[0] != 10 {
		t.Fatal("first value should be the same")
	}
	if s[1] != 20 {
		t.Fatal("second value should be the same")
	}
}

func TestSliceValueReader(t *testing.T) {
	s := []int64{10, 20}
	slice := gotagio.NewSliceValueReader(testInt64ValueReader)
	result := readValue(t, slice, s)
	if len(result) != 2 {
		t.Fatal("length should remain the same")
	}
	if result[0] != "10" {
		t.Fatal("first value should be the same")
	}
	if result[1] != "20" {
		t.Fatal("second value should be the same")
	}
}
