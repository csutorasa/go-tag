package gotagio_test

import (
	"fmt"
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestSetString(t *testing.T) {
	var s string
	writeValue(t, gotagio.WriteString, &s, "testValue")
	if s != "testValue" {
		t.Fatal("failed to set string value")
	}
}

func TestGetString(t *testing.T) {
	s := "testValue"
	result := readValue(t, gotagio.ReadString, s)
	if result != "testValue" {
		t.Fatal("failed to get string value")
	}
}

type StringerTest struct {
	s string
}

func (s StringerTest) String() string {
	return s.s
}

func TestSetStringer(t *testing.T) {
	var s string
	writeValue(t, gotagio.WriteStringer, &s, fmt.Stringer(StringerTest{s: "testValue"}))
	if s != "testValue" {
		t.Fatal("failed to set string value")
	}
}
