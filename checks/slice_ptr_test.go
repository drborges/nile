package checks_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"github.com/smartystreets/assertions/should"
	"github.com/drborges/nile/checks"
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
