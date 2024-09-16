package gotagio

import (
	"reflect"
	"strconv"

	"github.com/csutorasa/go-tags/gotag"
)

// [gotag.ValueWriterFunc] that parses the bool and sets the value to the [reflect.Value].
func WriteStrConvBool(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || v.Kind() != reflect.Bool {
		return false, nil
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.SetBool(b)
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the integer and sets the value to the [reflect.Value].
func WriteStrConvInt(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || !isInt(v) {
		return false, nil
	}
	kind := v.Type().Kind()
	var i int64
	var err error
	switch kind {
	case reflect.Int:
		i, err = strconv.ParseInt(s, 10, 0)
	case reflect.Int8:
		i, err = strconv.ParseInt(s, 10, 8)
	case reflect.Int16:
		i, err = strconv.ParseInt(s, 10, 16)
	case reflect.Int32:
		i, err = strconv.ParseInt(s, 10, 32)
	case reflect.Int64:
		i, err = strconv.ParseInt(s, 10, 64)
	}
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.SetInt(i)
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the unsigned integer and sets the value to the [reflect.Value].
func WriteStrConvUint(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || !isUint(v) {
		return false, nil
	}
	kind := v.Type().Kind()
	var i uint64
	var err error
	switch kind {
	case reflect.Uint:
		i, err = strconv.ParseUint(s, 10, 0)
	case reflect.Uint8:
		i, err = strconv.ParseUint(s, 10, 8)
	case reflect.Uint16:
		i, err = strconv.ParseUint(s, 10, 16)
	case reflect.Uint32:
		i, err = strconv.ParseUint(s, 10, 32)
	case reflect.Uint64:
		i, err = strconv.ParseUint(s, 10, 64)
	}
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.SetUint(i)
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the floating point number and sets the value to the [reflect.Value].
func WriteStrConvFloat(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || !isFloat(v) {
		return false, nil
	}
	kind := v.Type().Kind()
	var f float64
	var err error
	switch kind {
	case reflect.Float32:
		f, err = strconv.ParseFloat(s, 32)
	case reflect.Float64:
		f, err = strconv.ParseFloat(s, 64)
	}
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.SetFloat(f)
	return true, nil
}

// [gotag.ValueWriterFunc] that parses the complex number and sets the value to the [reflect.Value].
func WriteStrConvComplex(v reflect.Value, s string) (bool, error) {
	if !v.CanSet() || !isComplex(v) {
		return false, nil
	}
	kind := v.Type().Kind()
	var c complex128
	var err error
	switch kind {
	case reflect.Complex64:
		c, err = strconv.ParseComplex(s, 64)
	case reflect.Complex128:
		c, err = strconv.ParseComplex(s, 128)
	}
	if err != nil {
		return true, gotag.NewWriteValueError(s, v.Type(), err)
	}
	v.SetComplex(c)
	return true, nil
}

// A combined [gotag.ValueWriterFunc] for all [strconv] [gotag.ValueWriterFunc]s.
var WriteStrConv gotag.ValueWriterFunc[string] = gotag.NewFirstSupportedValueWriter(
	WriteStrConvBool,
	WriteStrConvInt,
	WriteStrConvUint,
	WriteStrConvFloat,
	WriteStrConvComplex,
)

// [gotag.ValueReaderFunc] that returns a formatted the bool from the [reflect.Value].
func ReadStrConvBool(v reflect.Value) (string, bool, error) {
	if v.Kind() != reflect.Bool {
		return "", false, nil
	}
	b := v.Bool()
	return strconv.FormatBool(b), true, nil
}

// [gotag.ValueReaderFunc] that returns a formatted the int from the [reflect.Value].
func ReadStrConvInt(v reflect.Value) (string, bool, error) {
	if !isInt(v) {
		return "", false, nil
	}
	i := v.Int()
	return strconv.FormatInt(i, 10), true, nil
}

// [gotag.ValueReaderFunc] that returns a formatted the unsigned int from the [reflect.Value].
func ReadStrConvUint(v reflect.Value) (string, bool, error) {
	if !isUint(v) {
		return "", false, nil
	}
	i := v.Uint()
	return strconv.FormatUint(i, 10), true, nil
}

// [gotag.ValueReaderFunc] that returns a formatted the floating point number from the [reflect.Value].
func ReadStrConvFloat(v reflect.Value) (string, bool, error) {
	if !isFloat(v) {
		return "", false, nil
	}
	f := v.Float()
	return strconv.FormatFloat(f, 'f', -1, int(v.Type().Size())*8), true, nil
}

// [gotag.ValueReaderFunc] that returns a formatted the complex number from the [reflect.Value].
func ReadStrConvComplex(v reflect.Value) (string, bool, error) {
	if !isComplex(v) {
		return "", false, nil
	}
	c := v.Complex()
	return strconv.FormatComplex(c, 'f', -1, int(v.Type().Size())*8), true, nil
}

// A combined [gotag.ValueReaderFunc] for all [strconv] [gotag.ValueReaderFunc]s.
var ReadStrConv gotag.ValueReaderFunc[string] = gotag.NewFirstSupportedValueReader(
	ReadStrConvBool,
	ReadStrConvInt,
	ReadStrConvUint,
	ReadStrConvFloat,
	ReadStrConvComplex,
)

func isInt(v reflect.Value) bool {
	fieldKind := v.Kind()
	return fieldKind == reflect.Int || fieldKind == reflect.Int8 || fieldKind == reflect.Int16 || fieldKind == reflect.Int32 || fieldKind == reflect.Int64
}

func isUint(v reflect.Value) bool {
	fieldKind := v.Kind()
	return fieldKind == reflect.Uint || fieldKind == reflect.Uint8 || fieldKind == reflect.Uint16 || fieldKind == reflect.Uint32 || fieldKind == reflect.Uint64
}

func isFloat(v reflect.Value) bool {
	fieldKind := v.Kind()
	return fieldKind == reflect.Float32 || fieldKind == reflect.Float64
}

func isComplex(v reflect.Value) bool {
	fieldKind := v.Kind()
	return fieldKind == reflect.Complex64 || fieldKind == reflect.Complex128
}
