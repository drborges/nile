package checks_test

import (
	"github.com/drborges/nile/checks"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestPtrCheck(t *testing.T) {
	Convey("It returns an error for non pointers", t, func() {
		values := []reflect.Value{
			reflect.ValueOf(1),
			reflect.ValueOf("a"),
			reflect.ValueOf(checks.ErrCheck{}),
			reflect.ValueOf([]int{1, 2, 3}),
		}

		for _, value := range values {
			err := checks.Ptr(value)

			So(reflect.ValueOf(err).Type(), should.Resemble, reflect.ValueOf(checks.ErrCheck{}).Type())
		}
	})

	Convey("It returns nil for pointers", t, func() {
		num := 1
		str := "a"

		values := []reflect.Value{
			reflect.ValueOf(&num),
			reflect.ValueOf(&str),
			reflect.ValueOf(&checks.ErrCheck{}),
			reflect.ValueOf(&[]int{1, 2, 3}),
		}

		for _, value := range values {
			err := checks.Ptr(value)

			So(err, should.BeNil)
		}
	})
}
