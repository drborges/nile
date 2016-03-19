package checks_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"github.com/smartystreets/assertions/should"
	"github.com/drborges/nile/checks"
)

func TestFlattenableCheck(t *testing.T) {
	Convey("It returns an error for non flattenables", t, func() {
		values := []reflect.Value{
			reflect.ValueOf(1),
			reflect.ValueOf("a"),
			reflect.ValueOf(checks.ErrCheck{}),
		}

		for _, value := range values {
			err := checks.Flattenable(value)

			So(reflect.ValueOf(err).Type(), should.Resemble, reflect.ValueOf(checks.ErrCheck{}).Type())
		}
	})

	Convey("It returns nil for flattenables", t, func() {
		arr := make([]int, 3)

		values := []reflect.Value{
			reflect.ValueOf(arr),
			reflect.ValueOf(&arr),
			reflect.ValueOf([]int{}),
			reflect.ValueOf(&[]int{}),
		}

		for _, value := range values {
			err := checks.Flattenable(value)

			So(err, should.BeNil)
		}
	})
}
