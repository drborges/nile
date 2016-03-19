package transformers_test

import (
	"errors"
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/stream"
	"github.com/drborges/nile/transformers"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestMapper(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And a mapper bound to the context", func() {
			addOne := func(data stream.T) stream.T { return data.(int) + 1 }
			addOneTransform := transformers.Mapper(addOne)(ctx)

			Convey("When a stream of numbers is produced", func() {
				r, w := stream.New(2)
				w <- 1
				w <- 2
				close(w)

				Convey("Then all items are mapped", func() {
					stream := addOneTransform(r)

					So(<-stream, should.Equal, 2)
					So(<-stream, should.Equal, 3)

					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done without errors", func() {
				ctx.Signal(context.Done)

				Convey("Then data is still mapped from the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					stream := addOneTransform(r)

					So(<-stream, should.Equal, 2)
					So(<-stream, should.Equal, 3)

					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When context is done with errors", func() {
				ctx.Signal(errors.New("Oops..."))

				Convey("Then data is no longer mappedfrom the upstream", func() {
					r, w := stream.New(2)
					w <- 1
					w <- 2
					close(w)

					stream := addOneTransform(r)
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done

					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})
		})
	})
}
