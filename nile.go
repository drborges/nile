package nile

import (
	"github.com/drborges/nile/stream"
)

// TODO integrate with golang net/context
type Context interface {
	Done() <-chan struct{}
	Err() error
	Signal(err error)
}

type Mapping func(data stream.T) stream.T
type Predicate func(data stream.T) bool

type Producer func(Context) Produce
type Consumer func(Context) Consume
type Transformer func(Context) Transform

type Produce func() (out stream.Readable)
type Transform func(in stream.Readable) (out stream.Readable)
type Consume func(in stream.Readable)

func Compose(a Transformer, b Transformer) Transformer {
	return func(ctx Context) Transform {
		return func(in stream.Readable) stream.Readable {
			return b(ctx)(a(ctx)(in))
		}
	}
}

func HasErrors(ctx Context) <-chan struct{} {
	sig := make(chan struct{})

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			close(sig)
		}
	default:
	}

	return sig
}
