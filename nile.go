package nile

import (
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/stream"
)

// TODO integrate with golang net/context
type Context interface {
	Done() <-chan struct{}
	Err() error
	Signal(err error)
}

type Producer func(Context) Produce
type Consumer func(Context) Consume
type Transformer func(Context) Transform
type Piper func(Context) Pipe
type Merger func(Context) Merge
type Forker func(Context) Fork
type Runner func() error

func (runner Runner) Run() error {
	return runner()
}

type Mapping func(data stream.T) stream.T
type Predicate func(data stream.T) bool

type Produce func() (out stream.Readable)
type Transform func(in stream.Readable) (out stream.Readable)
type Consume func(in stream.Readable)
type Pipe func(from stream.Readable, to stream.Writable)
type Merge func(in ...stream.Readable) (out stream.Readable)
type Fork func(in stream.Readable) (out1 stream.Readable, out2 stream.Readable)

type Pipeline struct {
	ctx        Context
	source     Produce
	transforms Transform
	sink       Consume
}

func From(producer Producer) Pipeline {
	ctx := context.New()
	return Pipeline{
		ctx:        ctx,
		source:     producer(ctx),
		transforms: noop,
	}
}

func (pipe Pipeline) Apply(transformer Transformer) Pipeline {
	pipe.transforms = Compose(pipe.transforms, transformer(pipe.ctx))
	return pipe
}

func (pipe Pipeline) Then(consumer Consumer) Runner {
	pipe.sink = consumer(pipe.ctx)
	return func() error {
		pipe.sink(pipe.transforms(pipe.source()))
		return pipe.ctx.Err()
	}
}

func noop(in stream.Readable) (out stream.Readable) {
	return in
}

func Compose(a Transform, b Transform) Transform {
	return func(in stream.Readable) stream.Readable {
		return b(a(in))
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
