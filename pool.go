package hzypool

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type pool struct {
	guard           sync.RWMutex
	workPool        chan *Worker
	jobs            chan *Job
	maxWorkNum      int
	maxWorkDuration time.Duration
	logger          Logger
	runningWorkNum  int
}

func New(sl ...Setter) (*pool, error) {
	p := &pool{}
	p.maxWorkNum = runtime.NumCPU() * 2
	for _, s := range sl {
		s(p)
	}

	if p.maxWorkNum < 1 {
		return nil, fmt.Errorf("must have at least one worker in the pool")
	}

	p.workPool = make(chan *Worker, p.maxWorkNum)
	p.jobs = make(chan *Job, 1)

	p.dispatch()

	return p, nil
}

func (p *pool) add(w *Worker) {
	p.guard.Lock()
	p.runningWorkNum++
	p.guard.Unlock()
}
func (p *pool) Submit(j *Job) {
	if j == nil {
		return
	}

	p.jobs <- j
}

func (p *pool) NewWorker() *Worker {
	w := newWorker(p)
	p.add(w)
	return w
}

func (p *pool) dispatch() {
	go func() {
		for j := range p.jobs {
			for {
				select {
				case w := <-p.workPool:
					w.submit(j)
				default:
					if p.runningWorkNum < p.maxWorkNum {
						p.NewWorker().submit(j)
					} else {
						w := <-p.workPool
						w.submit(j)
					}
				}
			}
		}
	}()
}
