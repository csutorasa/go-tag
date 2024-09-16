package gotagio

import (
	"io"
	"reflect"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that reads the [io.Reader] and sets the content to a byte array [reflect.Value].
func WriteByteArrayFromReader(v reflect.Value, r io.Reader) (bool, error) {
	if v.Kind() != reflect.Array || v.Type().Elem().Kind() != reflect.Uint8 {
		return false, nil
	}
	l := v.Type().Len()
	b, err := io.ReadAll(io.LimitReader(r, int64(l)))
	if err != nil {
		return true, gotag.NewWriteValueError(r, v.Type(), err)
	}
	reflect.Copy(v, reflect.ValueOf(b))
	return true, nil
}

// [gotag.ValueWriterFunc] that reads the [io.Reader] and sets the content to a byte slice [reflect.Value].
func WriteBytesFromReader(v reflect.Value, r io.Reader) (bool, error) {
	if !v.CanSet() || v.Kind() != reflect.Slice || v.Type().Elem().Kind() != reflect.Uint8 {
		return false, nil
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return true, gotag.NewWriteValueError(r, v.Type(), err)
	}
	v.Set(reflect.ValueOf(b))
	return true, nil
}

// [gotag.ValueWriterFunc] that passes the [io.Reader] to a [reflect.Value].
func WriteReaderFromReader[T io.Reader](v reflect.Value, r io.Reader) (bool, error) {
	if !v.CanSet() || v.Type() != reflect.TypeFor[T]() {
		return false, nil
	}
	v.Set(reflect.ValueOf(r))
	return true, nil
}

// [gotag.ValueWriterFunc] that reads the [io.Reader] and sets the content to a string [reflect.Value].
func WriteStringFromReader(v reflect.Value, r io.Reader) (bool, error) {
	if !v.CanSet() || v.Kind() != reflect.String {
		return false, nil
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return true, gotag.NewWriteValueError(r, v.Type(), err)
	}
	v.SetString(string(b))
	return true, nil
}

// A combined [gotag.ValueWriterFunc] for all [io.WriteFromReader] based [gotag.ValueWriterFunc]s.
var WriteFromReader gotag.ValueWriterFunc[io.Reader] = gotag.NewValueWriters(
	WriteByteArrayFromReader,
	WriteBytesFromReader,
	WriteReaderFromReader[io.Reader],
	WriteStringFromReader,
)
