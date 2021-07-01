package hzypool

import "context"

type fn func(ctx context.Context, arg interface{}) error

type job struct {
	Fn  fn
	Arg interface{}
}

type worker struct {
	job chan *job
}

func (w *worker) submit(j *job) {
	if j == nil {
		return
	}
	w.job <- j
}

func (w *worker) do(j *job) error {
	return j.Fn(context.Background(), j.Arg)
}
