package consumers

import (
	"reflect"
	"github.com/drborges/nile"
	"github.com/drborges/nile/checks"
	"github.com/drborges/nile/stream"
)

func Collector(dst interface{}) nile.Consumer {
	val := reflect.ValueOf(dst)
	if err := checks.SlicePtr(val); err != nil {
		panic(err)
	}

	return func(ctx nile.Context) nile.Consume {
		return func(in stream.Readable) {
			for data := range in {
				select {
				case <-nile.HasErrors(ctx):
					return
				default:
					slice := val.Elem()
					slice.Set(reflect.Append(slice, reflect.ValueOf(data)))
				}
			}
		}
	}
}
