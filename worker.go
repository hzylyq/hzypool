package hzypool

import "context"

type Worker struct {
	p    *pool
	Jobs chan *Job
}

func newWorker(p *pool) *Worker {
	w := &Worker{
		p:    p,
		Jobs: make(chan *Job, 1),
	}
	w.schedule()
	return w
}

func (w *Worker) submit(j *Job) {
	if j == nil {
		return
	}
	w.Jobs <- j
}

func (w *Worker) schedule() {
	go func() {
		for j := range w.Jobs {
			w.exec(j)
		}
	}()
}

func (w *Worker) exec(j *Job) {
	ctx, cancel := context.WithTimeout(context.Background(), w.p.maxWorkDuration)
	defer func() {
		if r := recover(); r != nil {
			w.p.logger.Printf("%s", r)
		}
		cancel()
		w.p.workPool <- w
	}()
	if err := j.Fn(ctx, j.Arg); err != nil {
		w.p.logger.Printf("%s", err)
	}
}
