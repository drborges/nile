package nile_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/transformers"
	"github.com/drborges/nile/stream"
	"github.com/smartystreets/assertions/should"
	"github.com/drborges/nile"
)

func TestCompose(t *testing.T) {
	Convey("Given a new context", t, func() {
		ctx := context.New()

		Convey("And two filters bound to the context", func() {
			evens := func(data stream.T) bool { return data.(int) % 2 == 0 }
			gt4 := func(data stream.T) bool { return data.(int) > 4 }

			evensOnlyFilter := transformers.Filter(evens)(ctx)
			gt4Filter := transformers.Filter(gt4)(ctx)

			Convey("When a stream of numbers is filtered by the composition of these filters", func() {
				r, w := stream.New(4)
				w <- 1; w <- 2; w <- 6; w <- 7
				close(w)

				filter := nile.Compose(evensOnlyFilter, gt4Filter)

				Convey("Then only filtered items should be sent downstream", func() {
					stream := filter(r)

					So(<-stream, should.Equal, 6)
					data, more := <-stream
					So(data, should.BeNil)
					So(more, should.BeFalse)
				})
			})
		})
	})
}
