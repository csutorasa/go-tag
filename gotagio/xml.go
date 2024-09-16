package gotagio

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that reads the reader, parses the xml and sets the value to the [reflect.Value].
func WriteXmlReader(v reflect.Value, r io.Reader) (bool, error) {
	if !v.CanAddr() {
		return false, nil
	}
	addr := v.Addr()
	if !addr.CanInterface() {
		return false, nil
	}
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(addr.Interface())
	if err != nil {
		return true, gotag.NewWriteValueError(r, v.Type(), err)
	}
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the xml and sets the value to the [reflect.Value].
func WriteXmlString(v reflect.Value, s string) (bool, error) {
	if !v.CanAddr() {
		return false, nil
	}
	addr := v.Addr()
	if !addr.CanInterface() {
		return false, nil
	}
	err := xml.Unmarshal([]byte(s), addr.Interface())
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the xml and sets the value to the [reflect.Value].
func WriteXmlBytes(v reflect.Value, b []byte) (bool, error) {
	if !v.CanAddr() {
		return false, nil
	}
	addr := v.Addr()
	if !addr.CanInterface() {
		return false, nil
	}
	err := xml.Unmarshal(b, addr.Interface())
	if err != nil {
		return true, gotag.NewWriteValueError(b, v.Type(), err)
	}
	return true, nil
}

// [gotag.ValueReaderFunc] that returns the stringified the xml from the [reflect.Value].
func ReadXmlString(v reflect.Value) (string, bool, error) {
	if !v.CanInterface() {
		return "", false, nil
	}
	b, err := xml.Marshal(v.Interface())
	if err != nil {
		return "", true, gotag.NewReadValueError[string](v.Type(), err)
	}
	return string(b), true, nil
}

// [gotag.ValueReaderFunc] that parses the xml and sets the value to the [reflect.Value].
func ReadXmlBytes(v reflect.Value) ([]byte, bool, error) {
	if !v.CanInterface() {
		return nil, false, nil
	}
	b, err := xml.Marshal(v.Interface())
	if err != nil {
		return nil, true, gotag.NewReadValueError[[]byte](v.Type(), err)
	}
	return b, true, nil
}
