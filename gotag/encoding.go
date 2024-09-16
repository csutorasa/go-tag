package gotag

import (
	"reflect"
	"strings"
)

// Decoder that creates new instaces of structs that have struct tags.
type StructDecoder[T any, S any] interface {
	// Create a new instance based on the source.
	Decode(s S) (T, error)
}

// Creates a default [gotag.StructDecoder].
func NewDecoder[T any, S any](writers ...StructTagValueWriter[S]) StructDecoder[T, S] {
	return &defaultStructDecoder[T, S]{
		writers: writers,
	}
}

type defaultStructDecoder[T any, R any] struct {
	writers []StructTagValueWriter[R]
}

// Decode implements StructDecoder.
func (d *defaultStructDecoder[T, R]) Decode(r R) (T, error) {
	var result T
	err := forEach(&result, d.writers, func(w StructTagValueWriter[R], cmd *StructTagCommand) error {
		return w.Write(cmd, r)
	})
	if err != nil {
		var t T
		return t, err
	}
	return result, nil
}

// Decoder that creates new instaces of structs that have struct tags.
type StructEncoder[T any, R any] interface {
	// Create a new instance based on the source.
	Encode(t T, defaultValue R) (R, error)
}

type StructEncoderCombiner[S any, R any] func(S, R) (R, error)

// Creates a default [gotag.StructEncoder].
func NewEncoder[T any, R any](readers ...StructTagValueReader[R]) StructEncoder[T, R] {
	return &defaultStructEncoder[T, R]{
		readers: readers,
	}
}

type defaultStructEncoder[T any, R any] struct {
	readers []StructTagValueReader[R]
}

// Encode implements StructEncoder.
func (e *defaultStructEncoder[T, R]) Encode(t T, defaultValue R) (R, error) {
	result := defaultValue
	err := forEach(&t, e.readers, func(r StructTagValueReader[R], cmd *StructTagCommand) error {
		res, err := r.Read(cmd, result)
		if err != nil {
			return err
		}
		result = res
		return nil
	})
	if err != nil {
		var r R
		return r, err
	}
	return result, nil
}

func forEach[T any, H StructTagHandler](t *T, handlers []H, f func(H, *StructTagCommand) error) error {
	value := reflect.ValueOf(t).Elem()
	structType := value.Type()
	if structType.Kind() != reflect.Struct {
		return ErrNotAStruct
	}
	for i := 0; i < structType.NumField(); i++ {
		structField := structType.Field(i)
		resultFieldValue := value.Field(i)
		for _, h := range handlers {
			tagValue, ok := structField.Tag.Lookup(h.Tag())
			if ok {
				cmd := &StructTagCommand{
					tag:         h.Tag(),
					tagValues:   getTags(tagValue),
					structType:  structType,
					structField: structField,
					fieldValue:  resultFieldValue,
				}
				err := f(h, cmd)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func getTags(s string) []string {
	tags := []string{}
	for _, tag := range strings.Split(s, ",") {
		t := strings.TrimSpace(tag)
		if t != "" {
			tags = append(tags, t)
		}
	}
	return tags
}
