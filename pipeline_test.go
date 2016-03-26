package nile_test

import (
	"github.com/drborges/nile"
	. "github.com/drborges/nile/consumers"
	"github.com/drborges/nile/context"
	. "github.com/drborges/nile/producers"
	"github.com/drborges/nile/stream"
	. "github.com/drborges/nile/transformers"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPipeline(t *testing.T) {
	Convey("From(Range) -> Apply(Filter) -> Then(Collect) -> Run()", t, func() {
		var evens []int
		evensOnly := func(data stream.T) bool { return data.(int)%2 == 0 }

		err := nile.From(Range(0, 5)).Apply(Filter(evensOnly)).Then(Collect(&evens)).Run()

		So(err, should.BeNil)
		So(evens, should.Contain, 0)
		So(evens, should.Contain, 2)
		So(evens, should.Contain, 4)
		So(len(evens), should.Equal, 3)
	})

	Convey("From(Range) -> Then(Collect) -> Run()", t, func() {
		var evens []int

		err := nile.From(Range(0, 3)).Then(Collect(&evens)).Run()

		So(err, should.BeNil)
		So(evens, should.Contain, 0)
		So(evens, should.Contain, 1)
		So(evens, should.Contain, 2)
		So(len(evens), should.Equal, 3)
	})

	Convey("From(Range) -> Then(Collect) -> RunWith(ctx)", t, func() {
		var evens []int

		ctx := context.New()

		err := nile.From(Range(0, 3)).Then(Collect(&evens)).RunWith(ctx)

		So(err, should.BeNil)
		So(evens, should.Contain, 0)
		So(evens, should.Contain, 1)
		So(evens, should.Contain, 2)
		So(len(evens), should.Equal, 3)
	})

	Convey("From(Range) -> Then(PipeTo(w)) -> RunWith(ctx)", t, func() {
		r, w := stream.New(3)

		err := nile.From(Range(0, 3)).Then(PipeTo(w)).Run()

		So(err, should.BeNil)
		So(<-r, should.Equal, 0)
		So(<-r, should.Equal, 1)
		So(<-r, should.Equal, 2)

		data, more := <-r
		So(data, should.BeNil)
		So(more, should.BeFalse)
	})

	Convey("From(Range) -> Split() -> RunWith(ctx)", t, func() {
		p1, p2, runner := nile.From(Range(0, 3)).Split()

		runner1 := p1.Apply(Filter(evensOnly)).Then(Collect(&evens))
		runner2 := p2.Apply(Filter(oddsOnly)).Then(Collect(&odds))

		err := nile.RunAll(runner, runner1, runner2)

		So(err, should.BeNil)
		So(<-r, should.Equal, 0)
		So(<-r, should.Equal, 1)
		So(<-r, should.Equal, 2)

		data, more := <-r
		So(data, should.BeNil)
		So(more, should.BeFalse)
	})
}
