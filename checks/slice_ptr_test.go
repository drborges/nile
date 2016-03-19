package checks_test

import (
	"github.com/drborges/nile/checks"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestSlicePtrCheck(t *testing.T) {
	Convey("It returns an error for non pointers", t, func() {
		arr := make([]int, 3)
		values := []reflect.Value{
			reflect.ValueOf(1),
			reflect.ValueOf("a"),
			reflect.ValueOf(checks.ErrCheck{}),
			reflect.ValueOf(arr),
			reflect.ValueOf([]int{}),
		}

		for _, value := range values {
			err := checks.SlicePtr(value)

			So(reflect.ValueOf(err).Type(), should.Resemble, reflect.ValueOf(checks.ErrCheck{}).Type())
		}
	})

	Convey("It returns nil for slice pointers", t, func() {
		arr := make([]int, 3)

		values := []reflect.Value{
			reflect.ValueOf(&arr),
			reflect.ValueOf(&[]int{}),
		}

		for _, value := range values {
			err := checks.SlicePtr(value)

			So(err, should.BeNil)
		}
	})
}
