package checks

import (
	"fmt"
	"reflect"
)

type ErrCheck struct {
	expected reflect.Kind
	actual   reflect.Kind
}

func (err ErrCheck) Error() string {
	return fmt.Sprintf("Invalid Type. Expected %v, got %v", err.expected, err.actual)
}
