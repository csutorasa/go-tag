package gotagio_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestReaderBytesArray(t *testing.T) {
	var a [9]byte
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteByteArrayFromReader, &a, io.Reader(r))
	if string(a[:]) != "testValue" {
		t.Fatal("value should be set")
	}
}

func TestReaderBytesArrayShort(t *testing.T) {
	var a [4]byte
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteByteArrayFromReader, &a, io.Reader(r))
	if string(a[:]) != "test" {
		t.Fatal("value should be set")
	}
}

func TestReaderBytes(t *testing.T) {
	var b []byte
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteBytesFromReader, &b, io.Reader(r))
	if string(b) != "testValue" {
		t.Fatal("value should be set")
	}
}

func TestReaderReader(t *testing.T) {
	var s io.Reader
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteReaderFromReader[io.Reader], &s, io.Reader(r))
	if s != r {
		t.Fatal("readers should be the same")
	}
}

func TestReaderReaderGeneric(t *testing.T) {
	var s *bytes.Reader
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteReaderFromReader[*bytes.Reader], &s, io.Reader(r))
	if s != r {
		t.Fatal("readers should be the same")
	}
}

func TestReaderString(t *testing.T) {
	var s string
	r := bytes.NewReader([]byte("testValue"))
	writeValue(t, gotagio.WriteStringFromReader, &s, io.Reader(r))
	if s != "testValue" {
		t.Fatal("value should be set")
	}
}
