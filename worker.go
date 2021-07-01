package hzypool

import "context"

type fn func(ctx context.Context, arg interface{}) error

type Worker struct {
	Fn  fn
	Arg interface{}
}

func (w *Worker) submit() {

}

func (w *Worker) do() error {
	return w.Fn(context.Background(), w.Arg)
}
