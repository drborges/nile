package checks

import (
	"reflect"
)

func Flattenable(value reflect.Value) error {
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		return nil
	}

	if value.Kind() == reflect.Ptr && (value.Elem().Kind() == reflect.Slice || value.Elem().Kind() == reflect.Array) {
		return nil
	}

	return ErrCheck{
		expected: reflect.Slice,
		actual:   value.Kind(),
	}
}
