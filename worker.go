package hzypool

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

}
