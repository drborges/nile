package checks

import (
	"reflect"
)

func SlicePtr(value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return ErrCheck{
			expected: reflect.Ptr,
			actual:   value.Kind(),
		}
	}

	if value.Elem().Kind() != reflect.Slice && value.Elem().Kind() != reflect.Array {
		return ErrCheck{
			expected: reflect.Slice,
			actual:   value.Elem().Kind(),
		}
	}

	return nil
}
