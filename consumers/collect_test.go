package consumers_test

import (
	"errors"
	"github.com/drborges/nile/consumers"
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/stream"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCollect(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And a collector bound to the context", func() {
			var numbers []int
			collect := consumers.Collect(&numbers)(ctx)

			Convey("When a stream of numbers is collected by the collector", func() {
				r, w := stream.New(2)
				w <- 1
				w <- 2
				close(w)

				collect(r)

				Convey("Then all numbers are collected", func() {
					So(numbers, should.Contain, 1)
					So(numbers, should.Contain, 2)
					So(len(numbers), should.Equal, 2)
				})
			})

			Convey("When context is done without errors", func() {
				ctx.Signal(context.Done)

				Convey("Then data is still collected from the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					collect(r)

					So(numbers, should.Contain, 1)
					So(numbers, should.Contain, 2)
					So(len(numbers), should.Equal, 2)
				})
			})

			Convey("When context is done with errors", func() {
				ctx.Signal(errors.New("Oops..."))

				Convey("Then data is no longer collected from the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					collect(r)
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done

					So(numbers, should.BeEmpty)
				})
			})
		})
	})
}
