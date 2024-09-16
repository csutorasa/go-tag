package gotagio_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestArrayValueWriter(t *testing.T) {
	var s [2]int64
	array := gotagio.NewArrayValueWriter(testInt64ValueWriter)
	writeValue(t, array, &s, []string{"10", "20"})
	if s[0] != 10 {
		t.Fatal("first value should be the same")
	}
	if s[1] != 20 {
		t.Fatal("second value should be the same")
	}
}

func TestArrayValueWriterShort(t *testing.T) {
	var s [1]int64
	array := gotagio.NewArrayValueWriter(testInt64ValueWriter)
	writeValue(t, array, &s, []string{"10", "20"})
	if s[0] != 10 {
		t.Fatal("first value should be the same")
	}
}

func TestSliceValueReaderArray(t *testing.T) {
	s := [2]int64{10, 20}
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
