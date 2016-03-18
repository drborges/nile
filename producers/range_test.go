package producers_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/producers"
	"github.com/smartystreets/assertions/should"
	"errors"
	"time"
)

func TestRange(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And a range producer bound to the context", func() {
			produce := producers.Range(0, 3)(ctx)

			Convey("When data is produced", func() {
				stream := produce()

				Convey("Then data can be consumed from the stream until it is done", func() {
					So(<-stream, should.Equal, 0)
					So(<-stream, should.Equal, 1)
					So(<-stream, should.Equal, 2)

					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When the context is done without errors", func() {
				ctx.Signal(context.Done)

				Convey("Then no data is produced", func() {
					stream := produce()
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done
					data, more := <-stream

					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})

			Convey("When the context is done with errors", func() {
				ctx.Signal(errors.New("Oops..."))

				Convey("Then no data is produced", func() {
					stream := produce()
					time.Sleep(100 * time.Millisecond) // give producer some time to realize the context is done
					data, more := <-stream

					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})
		})
	})
}