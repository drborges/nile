package producers

import (
	"github.com/drborges/nile"
	"github.com/drborges/nile/stream"
)

func Range(from, to int) nile.Producer {
	return func(ctx nile.Context) nile.Produce {
		return func() stream.Readable {
			r, w := stream.New(0)

			go func() {
				defer close(w)

				for i := 0; i < to; i++ {
					select {
					case <-ctx.Done():
						return
					case w <- i:
					}
				}
			}()

			return r
		}
	}
}
