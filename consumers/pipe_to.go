package consumers

import (
	"github.com/drborges/nile"
	"github.com/drborges/nile/stream"
)

func PipeTo(w stream.Writable) nile.Consumer {
	return func(ctx nile.Context) nile.Consume {
		return func(in stream.Readable) {
			defer close(w)

			for data := range in {
				select {
				case <-nile.HasErrors(ctx):
					return
				case w <- data:
				}
			}
		}
	}
}
