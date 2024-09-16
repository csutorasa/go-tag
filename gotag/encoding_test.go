package gotag_test

import (
	"reflect"
	"testing"

	"github.com/csutorasa/go-tags/gotag"
)

type mockStructTagValueWriter struct {
	tag string
	f   func(command *gotag.StructTagCommand, r string) error
}

func (w *mockStructTagValueWriter) Tag() string {
	return w.tag
}

func (w *mockStructTagValueWriter) Write(cmd *gotag.StructTagCommand, r string) error {
	return w.f(cmd, r)
}

func testStringValueWriter(v reflect.Value, s string) (bool, error) {
	v.SetString(s)
	return true, nil
}

type StructWithTags struct {
	S string `test:"value"`
}

func TestDecode(t *testing.T) {
	ok := false
	var w gotag.StructTagValueWriter[string] = &mockStructTagValueWriter{
		tag: "test",
		f: func(cmd *gotag.StructTagCommand, s string) error {
			ok = cmd.Tag() == "test" && len(cmd.TagValues()) == 1 && cmd.TagValues()[0] == "value"
			var c gotag.ValueWriterFunc[string] = testStringValueWriter
			return c.WriteValue(cmd, s)
		},
	}
	d := gotag.NewDecoder[StructWithTags](w)
	result, err := d.Decode("otherString")
	if err != nil {
		t.Fatalf("decoder failed with %s", err.Error())
	}
	if !ok {
		t.Fatal("writer was not called correctly")
	}
	if result.S != "otherString" {
		t.Fatal("field was not set correctly")
	}
}

type mockStructTagValueReader struct {
	tag string
	f   func(*gotag.StructTagCommand, *ResultStruct) (*ResultStruct, error)
}

func (w *mockStructTagValueReader) Tag() string {
	return w.tag
}

func (w *mockStructTagValueReader) Read(cmd *gotag.StructTagCommand, result *ResultStruct) (*ResultStruct, error) {
	return w.f(cmd, result)
}

func testStringValueReader(v reflect.Value) (string, bool, error) {
	return v.String(), true, nil
}

type ResultStruct struct {
	S []string
}

func TestEncode(t *testing.T) {
	ok := false
	var r gotag.StructTagValueReader[*ResultStruct] = &mockStructTagValueReader{
		tag: "test",
		f: func(cmd *gotag.StructTagCommand, result *ResultStruct) (*ResultStruct, error) {
			ok = cmd.Tag() == "test" && len(cmd.TagValues()) == 1 && cmd.TagValues()[0] == "value"
			var c gotag.ValueReaderFunc[string] = testStringValueReader
			s, err := c.ReadValue(cmd)
			if err != nil {
				return nil, err
			}
			result.S = append(result.S, s)
			return result, nil
		},
	}
	e := gotag.NewEncoder[StructWithTags](r)
	result, err := e.Encode(StructWithTags{S: "otherString"}, &ResultStruct{S: []string{}})
	if err != nil {
		t.Fatalf("decoder failed with %s", err.Error())
	}
	if !ok {
		t.Fatal("writer was not called correctly")
	}
	if len(result.S) != 1 {
		t.Fatal("field was not set correctly")
	}
	if result.S[0] != "otherString" {
		t.Fatal("field was not set correctly")
	}
}
