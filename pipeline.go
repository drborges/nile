package nile

import (
	"github.com/drborges/nile/context"
	"github.com/drborges/nile/stream"
)

type Runner func(ctx Context) error

func (runner Runner) Run() error {
	return runner(context.New())
}

func (runner Runner) RunWith(ctx Context) error {
	return runner(ctx)
}

type Pipeline struct {
	source      Producer
	transformer Transformer
	sink        Consumer
}

func From(producer Producer) Pipeline {
	return Pipeline{
		source:      producer,
		transformer: noop,
	}
}

func (pipe Pipeline) Apply(transformer Transformer) Pipeline {
	pipe.transformer = Compose(pipe.transformer, transformer)
	return pipe
}

func (pipe Pipeline) Split() (Pipeline, Pipeline) {
	return Pipeline{}, Pipeline{}
}

func (pipe Pipeline) Then(consumer Consumer) Runner {
	pipe.sink = consumer
	return func(ctx Context) error {
		pipe.sink(ctx)(pipe.transformer(ctx)(pipe.source(ctx)()))
		return ctx.Err()
	}
}

func noop(ctx Context) Transform {
	return func(in stream.Readable) (out stream.Readable) {
		return in
	}
}
