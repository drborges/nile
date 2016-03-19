package transformers

import (
	"github.com/drborges/nile"
	"github.com/drborges/nile/stream"
)

func Filter(fn nile.Predicate) nile.Transformer {
	return func(ctx nile.Context) nile.Transform {
		return func(in stream.Readable) stream.Readable {
			r, w := stream.New(0)

			go func() {
				defer close(w)

				for data := range in {
					if fn(data) {
						select {
						case <-nile.HasErrors(ctx):
							return
						case w <- data:
						}
					}
				}
			}()

			return r
		}
	}
}
