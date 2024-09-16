package gotagio_test

import (
	"testing"

	"github.com/csutorasa/go-tags/gotagio"
)

func TestStrConvParseBool(t *testing.T) {
	var b bool
	writeValue(t, gotagio.WriteStrConvBool, &b, "true")
	if !b {
		t.Fatal("failed to set bool value")
	}
	writeValue(t, gotagio.WriteStrConvBool, &b, "false")
	if b {
		t.Fatal("failed to set bool value")
	}
}

func TestStrConvParseInt(t *testing.T) {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	writeValue(t, gotagio.WriteStrConvInt, &i, "42")
	if i != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvInt, &i8, "42")
	if i8 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvInt, &i16, "42")
	if i16 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvInt, &i32, "42")
	if i32 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvInt, &i64, "42")
	if i64 != 42 {
		t.Fatal("failed to set int value")
	}
}

func TestStrConvParseUint(t *testing.T) {
	var i uint
	var i8 uint8
	var i16 uint16
	var i32 uint32
	var i64 uint64
	writeValue(t, gotagio.WriteStrConvUint, &i, "42")
	if i != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvUint, &i8, "42")
	if i8 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvUint, &i16, "42")
	if i16 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvUint, &i32, "42")
	if i32 != 42 {
		t.Fatal("failed to set int value")
	}
	writeValue(t, gotagio.WriteStrConvUint, &i64, "42")
	if i64 != 42 {
		t.Fatal("failed to set int value")
	}
}

func TestStrConvParseFloat(t *testing.T) {
	var f32 float32
	var f64 float64
	writeValue(t, gotagio.WriteStrConvFloat, &f32, "42")
	if f32 != 42 {
		t.Fatal("failed to set float value")
	}
	writeValue(t, gotagio.WriteStrConvFloat, &f64, "42")
	if f64 != 42 {
		t.Fatal("failed to set float value")
	}
}

func TestStrConvParseComplex(t *testing.T) {
	var c64 complex64
	var c128 complex128
	writeValue(t, gotagio.WriteStrConvComplex, &c64, "42+1i")
	if real(c64) != 42 {
		t.Fatal("failed to set complex value")
	}
	if imag(c64) != 1 {
		t.Fatal("failed to set complex value")
	}
	writeValue(t, gotagio.WriteStrConvComplex, &c128, "42+1i")
	if real(c128) != 42 {
		t.Fatal("failed to set complex value")
	}
	if imag(c128) != 1 {
		t.Fatal("failed to set complex value")
	}
}

func TestStrConvFormatBool(t *testing.T) {
	s := readValue(t, gotagio.ReadStrConvBool, true)
	if s != "true" {
		t.Fatal("failed to format bool value")
	}
	s = readValue(t, gotagio.ReadStrConvBool, false)
	if s != "false" {
		t.Fatal("failed to format bool value")
	}
}

func TestStrConvFormatInt(t *testing.T) {
	s := readValue(t, gotagio.ReadStrConvInt, int(42))
	if s != "42" {
		t.Fatal("failed to format int value")
	}
	s = readValue(t, gotagio.ReadStrConvInt, int8(42))
	if s != "42" {
		t.Fatal("failed to format int value")
	}
	s = readValue(t, gotagio.ReadStrConvInt, int16(42))
	if s != "42" {
		t.Fatal("failed to format int value")
	}
	s = readValue(t, gotagio.ReadStrConvInt, int32(42))
	if s != "42" {
		t.Fatal("failed to format int value")
	}
	s = readValue(t, gotagio.ReadStrConvInt, int64(42))
	if s != "42" {
		t.Fatal("failed to format int value")
	}
}

func TestStrConvFormatUint(t *testing.T) {
	s := readValue(t, gotagio.ReadStrConvUint, uint(42))
	if s != "42" {
		t.Fatal("failed to format uint value")
	}
	s = readValue(t, gotagio.ReadStrConvUint, uint8(42))
	if s != "42" {
		t.Fatal("failed to format uint value")
	}
	s = readValue(t, gotagio.ReadStrConvUint, uint16(42))
	if s != "42" {
		t.Fatal("failed to format uint value")
	}
	s = readValue(t, gotagio.ReadStrConvUint, uint32(42))
	if s != "42" {
		t.Fatal("failed to format uint value")
	}
	s = readValue(t, gotagio.ReadStrConvUint, uint64(42))
	if s != "42" {
		t.Fatal("failed to format uint value")
	}
}

func TestStrConvFormatFloat(t *testing.T) {
	s := readValue(t, gotagio.ReadStrConvFloat, float32(42))
	if s != "42" {
		t.Fatal("failed to format float value")
	}
	s = readValue(t, gotagio.ReadStrConvFloat, float64(42))
	if s != "42" {
		t.Fatal("failed to format float value")
	}
}

func TestStrConvFormatComplex(t *testing.T) {
	s := readValue(t, gotagio.ReadStrConvComplex, 42+1i)
	if s != "(42+1i)" {
		t.Fatal("failed to format complex value")
	}
	s = readValue(t, gotagio.ReadStrConvComplex, 42+1i)
	if s != "(42+1i)" {
		t.Fatal("failed to format complex value")
	}
}
