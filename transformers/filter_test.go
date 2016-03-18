package transformers_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/transformers"
	"github.com/drborges/nile/stream"
	"github.com/smartystreets/assertions/should"
	"errors"
	"time"
)

func TestFilter(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And a filter bound to the context", func() {
			evensOnly := func(data stream.T) bool { return data.(int) % 2 == 0 }
			filter := transformers.Filter(evensOnly)(ctx)

			Convey("When a stream of numbers is filtered", func() {
				r, w := stream.New(2)
				w <- 1; w <- 2
				close(w)

				Convey("Then only filtered items should be sent downstream", func() {
					stream := filter(r)

					So(<-stream, should.Equal, 2)
					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done without errors", func() {
				ctx.Signal(context.Done)

				Convey("Then data is still filtered from the upstream", func() {
					r, w := stream.New(2)
					w <- 1; w <- 2
					close(w)

					stream := filter(r)

					So(<-stream, should.Equal, 2)
					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done with errors", func() {
				ctx.Signal(errors.New("Oops..."))

				Convey("Then data is still filtered from the upstream", func() {
					r, w := stream.New(2)
					w <- 1; w <- 2
					close(w)

					stream := filter(r)
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done

					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})
		})
	})
}
