package gotag_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csutorasa/go-tags/gotag"
)

type mockValueWriter struct {
	supports bool
	err      error
	called   int
}

func (w *mockValueWriter) Write(v reflect.Value, s string) (bool, error) {
	w.called++
	return w.supports, w.err
}

func TestNewFirstSupportedValueWriter(t *testing.T) {
	unsupported := &mockValueWriter{supports: false, err: nil, called: 0}
	fails := &mockValueWriter{supports: true, err: fmt.Errorf(""), called: 0}
	succeed := &mockValueWriter{supports: true, err: nil, called: 0}
	w := gotag.NewFirstSupportedValueWriter(
		unsupported.Write,
		fails.Write,
		succeed.Write,
	)
	supports, err := w(reflect.Value{}, "")
	if !supports {
		t.Fatal("it should be supported")
	}
	if err == nil {
		t.Fatal("it should fail")
	}
	if unsupported.called != 1 {
		t.Fatal("it should be called")
	}
	if fails.called != 1 {
		t.Fatal("it should be called")
	}
	if succeed.called != 0 {
		t.Fatal("it should not be called")
	}
}

func TestNewFirstSucceedValueWriter(t *testing.T) {
	unsupported := &mockValueWriter{supports: false, err: nil, called: 0}
	fails := &mockValueWriter{supports: true, err: fmt.Errorf(""), called: 0}
	succeed := &mockValueWriter{supports: true, err: nil, called: 0}
	w := gotag.NewFirstSucceedValueWriter(
		unsupported.Write,
		fails.Write,
		succeed.Write,
	)
	supports, err := w(reflect.Value{}, "")
	if !supports {
		t.Fatal("it should be supported")
	}
	if err != nil {
		t.Fatal("it should not fail")
	}
	if unsupported.called != 1 {
		t.Fatal("it should be called")
	}
	if fails.called != 1 {
		t.Fatal("it should be called")
	}
	if succeed.called != 1 {
		t.Fatal("it should be called")
	}
}

type mockValueReader struct {
	result   string
	supports bool
	err      error
	called   int
}

func (r *mockValueReader) Read(v reflect.Value) (string, bool, error) {
	r.called++
	return r.result, r.supports, r.err
}

func TestNewFirstSupportedValueReader(t *testing.T) {
	unsupported := &mockValueReader{result: "1", supports: false, err: nil, called: 0}
	fails := &mockValueReader{result: "2", supports: true, err: fmt.Errorf(""), called: 0}
	succeed := &mockValueReader{result: "3", supports: true, err: nil, called: 0}
	w := gotag.NewFirstSupportedValueReader(
		unsupported.Read,
		fails.Read,
		succeed.Read,
	)
	result, supports, err := w(reflect.Value{})
	if !supports {
		t.Fatal("it should be supported")
	}
	if err == nil {
		t.Fatal("it should fail")
	}
	if unsupported.called != 1 {
		t.Fatal("it should be called")
	}
	if fails.called != 1 {
		t.Fatal("it should be called")
	}
	if succeed.called != 0 {
		t.Fatal("it should not be called")
	}
	if result != "2" {
		t.Fatal("incorrect result")
	}
}

func TestNewFirstSucceedValueReader(t *testing.T) {
	unsupported := &mockValueReader{result: "1", supports: false, err: nil, called: 0}
	fails := &mockValueReader{result: "2", supports: true, err: fmt.Errorf(""), called: 0}
	succeed := &mockValueReader{result: "3", supports: true, err: nil, called: 0}
	w := gotag.NewFirstSucceedValueReader(
		unsupported.Read,
		fails.Read,
		succeed.Read,
	)
	result, supports, err := w(reflect.Value{})
	if !supports {
		t.Fatal("it should be supported")
	}
	if err != nil {
		t.Fatal("it should not fail")
	}
	if unsupported.called != 1 {
		t.Fatal("it should be called")
	}
	if fails.called != 1 {
		t.Fatal("it should be called")
	}
	if succeed.called != 1 {
		t.Fatal("it should be called")
	}
	if result != "3" {
		t.Fatal("incorrect result")
	}
}
