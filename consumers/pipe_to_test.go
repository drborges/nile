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

func TestPipeTo(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And a pipe bound to the context", func() {
			dstR, dstW := stream.New(0)
			pipe := consumers.PipeTo(dstW)(ctx)

			Convey("When a stream of numbers is piped", func() {
				r, w := stream.New(2)
				w <- 1
				w <- 2
				close(w)

				go pipe(r)

				Convey("Then all numbers are sent downstream", func() {
					So(<-dstR, should.Equal, 1)
					So(<-dstR, should.Equal, 2)
					data, more := <-dstR
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done without errors", func() {
				ctx.Signal(context.Done)

				Convey("Then data is still piped from the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					go pipe(r)

					So(<-dstR, should.Equal, 1)
					So(<-dstR, should.Equal, 2)
					data, more := <-dstR
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done with errors", func() {
				ctx.Signal(errors.New("Oops..."))

				Convey("Then data is no longer piped from the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					go pipe(r)
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done

					data, more := <-dstR
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})
		})
	})
}
