package transformers

import (
	"github.com/drborges/nile"
	"github.com/drborges/nile/stream"
)

func Mapper(fn nile.Mapping) nile.Transformer {
	return func(ctx nile.Context) nile.Transform {
		return func(in stream.Readable) stream.Readable {
			r, w := stream.New(0)

			go func() {
				defer close(w)

				for data := range in {
					select {
					case <-nile.HasErrors(ctx):
						return
					case w <- fn(data):
					}
				}
			}()

			return r
		}
	}
}
