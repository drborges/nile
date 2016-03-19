package nile

import "github.com/drborges/nile/stream"

type Pipeline struct {
	source      Producer
	transformer Transformer
	sink        Consumer
}

func From(producer Producer) Pipeline {
	noop := func(ctx Context) Transform {
		return func(in stream.Readable) (out stream.Readable) {
			return in
		}
	}

	return Pipeline{
		source:      producer,
		transformer: noop,
	}
}

func (pipe Pipeline) Apply(transformer Transformer) Pipeline {
	pipe.transformer = Compose(pipe.transformer, transformer)
	return pipe
}

func (pipe Pipeline) Then(consumer Consumer) Runner {
	pipe.sink = consumer
	return func(ctx Context) error {
		pipe.sink(ctx)(pipe.transformer(ctx)(pipe.source(ctx)()))
		return ctx.Err()
	}
}
