package gotaghttp_test

import (
	"io"
	"testing"

	"github.com/csutorasa/go-tags/gotag"
	"github.com/csutorasa/go-tags/gotaghttp"
)

type RequestBodyStringParams struct {
	Body string `requestBody:""`
}

func TestRequestBodyString(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyStringParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p RequestBodyStringParams) {
		if p.Body != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type RequestBodyByteArrayParams struct {
	Body [10]byte `requestBody:""`
}

func TestRequestBodyByteArray(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyByteArrayParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p RequestBodyByteArrayParams) {
		if string(p.Body[:4]) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type RequestBodyByteSliceParams struct {
	Body []byte `requestBody:""`
}

func TestRequestBodyByteSlice(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyByteSliceParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p RequestBodyByteSliceParams) {
		if string(p.Body) != "test" {
			t.Fail()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

type RequestBodyReaderParams struct {
	Body io.Reader `requestBody:""`
}

func TestRequestBodyReader(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyReaderParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p RequestBodyReaderParams) {
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

type RequestBodyReadCloserParams struct {
	Body io.Reader `requestBody:""`
}

func TestRequestBodyReadCloser(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyReadCloserParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "plain/text", []byte("test"), testCreator, func(p RequestBodyReadCloserParams) {
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

type RequestBodyJsonParams struct {
	Body map[string]any `requestBody:"json"`
}

func TestRequestBodyJson(t *testing.T) {
	testCreator := gotag.NewDecoder[RequestBodyJsonParams](gotaghttp.BodyWriter)
	err := doRequest("/", "/", "application/json", []byte("{\"test\": \"value\"}"), testCreator, func(p RequestBodyJsonParams) {
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
