package context

import "sync"

var Done = (error)(nil)

type Signal <-chan struct{}

type Context struct {
	sig   chan struct{}
	err   error
	mutex sync.Mutex
}

func New() *Context {
	return &Context{
		sig: make(chan struct{}),
	}
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.sig
}

func (ctx *Context) Err() error {
	return ctx.err
}

func (ctx *Context) Signal(err error) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()

	select {
	case <-ctx.sig:
		return
	default:
		close(ctx.sig)

		if err != nil {
			ctx.err = err
		}
	}
}
