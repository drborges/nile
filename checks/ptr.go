package checks

import (
	"reflect"
)

func Ptr(value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return ErrCheck{
			expected: reflect.Ptr,
			actual:   value.Kind(),
		}
	}

	return nil
}
