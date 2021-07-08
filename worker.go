package hzypool

import "context"

type Worker struct {
	p    *pool
	Jobs chan *Job
}

func newWorker() *Worker {
	return &Worker{}
}

func (w *Worker) submit(j *Job) {
	if j == nil {
		return
	}
	w.Jobs <- j
}

func (w *Worker) schedule() {
	go func() {
		for {
			select {
			case j := <-w.Jobs:
				w.exec(j)
			}
		}
	}()
}

func (w *Worker) exec(j *Job) {
	ctx, cancel := context.WithTimeout(context.Background(), w.p.maxWorkDuration)
	defer func() {
		if r := recover(); r != nil {

		}
		cancel()
		w.p.workPool <- w
	}()
	if err := j.Fn(ctx, j.Arg); err != nil {

	}
}
