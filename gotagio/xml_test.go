package gotagio_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

type XmlTest struct {
	Value string `xml:"value"`
}

var xmlString = "<XmlTest><value>testValue</value></XmlTest>"

func TestDecodeXml(t *testing.T) {
	var s XmlTest
	r := bytes.NewReader([]byte(xmlString))
	writeValue(t, gotagio.WriteXmlReader, &s, io.Reader(r))
	if s.Value != "testValue" {
		t.Fatal("tag content should be set")
	}
}

func TestDecodeXmlString(t *testing.T) {
	var s XmlTest
	writeValue(t, gotagio.WriteXmlString, &s, xmlString)
	if s.Value != "testValue" {
		t.Fatal("tag content should be set")
	}
}

func TestDecodeXmlBytes(t *testing.T) {
	var s XmlTest
	writeValue(t, gotagio.WriteXmlBytes, &s, []byte(xmlString))
	if s.Value != "testValue" {
		t.Fatal("tag content should be set")
	}
}

func TestEncodeXmlString(t *testing.T) {
	s := XmlTest{
		Value: "testValue",
	}
	result := readValue(t, gotagio.ReadXmlString, s)
	if result != xmlString {
		t.Fatal("wrong xml string result")
	}
}

func TestEncodeXmlBytes(t *testing.T) {
	s := XmlTest{
		Value: "testValue",
	}
	result := readValue(t, gotagio.ReadXmlBytes, s)
	if string(result) != xmlString {
		t.Fatal("wrong xml string result")
	}
}
