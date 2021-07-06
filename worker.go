package hzypool

import (
	"context"
	"log"
)

type fn func(ctx context.Context, arg interface{}) error

type Worker struct {
	Fn  fn
	Arg interface{}

	p *pool
}

func (w *Worker) submit() {

}

func (w *Worker) do() error {
	ctx, cancel := context.WithTimeout(context.Background(), w.p.maxWorkDuration)

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		cancel()
	}()

	return w.Fn(ctx, w.Arg)
}
