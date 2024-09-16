package gotagio_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

var jsonString = "{\"value\":\"testValue\"}"

func TestDecodeJson(t *testing.T) {
	var s map[string]any
	r := bytes.NewReader([]byte(jsonString))
	writeValue(t, gotagio.WriteJsonReader, &s, io.Reader(r))
	v, ok := s["value"]
	if !ok {
		t.Fatal("value should exist")
	}
	if v != "testValue" {
		t.Fatal("value should be set")
	}
}

func TestDecodeJsonString(t *testing.T) {
	var s map[string]any
	writeValue(t, gotagio.WriteJsonString, &s, jsonString)
	v, ok := s["value"]
	if !ok {
		t.Fatal("value should exist")
	}
	if v != "testValue" {
		t.Fatal("value should be set")
	}
}

func TestDecodeJsonBytes(t *testing.T) {
	var s map[string]any
	writeValue(t, gotagio.WriteJsonBytes, &s, []byte(jsonString))
	v, ok := s["value"]
	if !ok {
		t.Fatal("value should exist")
	}
	if v != "testValue" {
		t.Fatal("value should be set")
	}
}

func TestEncodeJsonString(t *testing.T) {
	s := map[string]any{
		"value": "testValue",
	}
	result := readValue(t, gotagio.ReadJsonString, s)
	if result != jsonString {
		t.Fatal("wrong json string result")
	}
}

func TestEncodeJsonBytes(t *testing.T) {
	s := map[string]any{
		"value": "testValue",
	}
	result := readValue(t, gotagio.ReadJsonBytes, s)
	if string(result) != jsonString {
		t.Fatal("wrong json string result")
	}
}
