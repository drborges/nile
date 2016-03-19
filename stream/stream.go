package stream

type T interface{}
type Readable <-chan T
type Writable chan<- T

var Empty = func() Readable {
	r, w := New(0)
	close(w)
	return r
}()

func New(cap int) (Readable, Writable) {
	ch := make(chan T, cap)
	return ch, ch
}
