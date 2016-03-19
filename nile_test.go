package nile_test

import (
	"github.com/drborges/nile"
	. "github.com/drborges/nile/consumers"
	. "github.com/drborges/nile/producers"
	"github.com/drborges/nile/stream"
	. "github.com/drborges/nile/transformers"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNile(t *testing.T) {
	Convey("From(Range) -> Apply(Filter) -> Then(Collect)", t, func() {
		var evens []int
		evensOnly := func(data stream.T) bool { return data.(int)%2 == 0 }

		err := nile.From(Range(0, 5)).Apply(Filter(evensOnly)).Then(Collect(&evens)).Run()

		So(err, should.BeNil)
		So(evens, should.Contain, 0)
		So(evens, should.Contain, 2)
		So(evens, should.Contain, 4)
		So(len(evens), should.Equal, 3)
	})
}
