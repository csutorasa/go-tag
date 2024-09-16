package gotaghttp_test

import (
	"io"
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type BodyStringParams struct {
	Body string `requestBody:""`
}

func TestBodyString(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyStringParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p BodyStringParams) {
		if p.Body != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type BodyByteArrayParams struct {
	Body [10]byte `requestBody:""`
}

func TestBodyByteArray(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyByteArrayParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p BodyByteArrayParams) {
		if string(p.Body[:4]) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type BodyByteSliceParams struct {
	Body []byte `requestBody:""`
}

func TestBodyByteSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyByteSliceParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p BodyByteSliceParams) {
		if string(p.Body) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type BodyReaderParams struct {
	Body io.Reader `requestBody:""`
}

func TestBodyReader(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyReaderParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p BodyReaderParams) {
		result, err := io.ReadAll(p.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(result) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type BodyReadCloserParams struct {
	Body io.Reader `requestBody:""`
}

func TestBodyReadCloser(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyReadCloserParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p BodyReadCloserParams) {
		result, err := io.ReadAll(p.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(result) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type BodyJsonParams struct {
	Body map[string]any `requestBody:"json"`
}

func TestBodyJson(t *testing.T) {
	testCreator := gotag.NewDecoder[BodyJsonParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "application/json", []byte("{\"test\": \"value\"}"), testCreator, func(p BodyJsonParams) {
		v, ok := p.Body["test"]
		if !ok {
			t.Fail()
		}
		if v != "value" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}
