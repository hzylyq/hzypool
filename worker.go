package hzypool

import "context"

type fn func(ctx context.Context, arg interface{}) error

type worker struct {
	Fn  fn
	Arg interface{}
}

func (w *worker) submit() {

}

func (w *worker) do() error {
	return w.Fn(context.Background(), w.Arg)
}
