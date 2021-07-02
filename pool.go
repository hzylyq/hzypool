package hzypool

import (
	"fmt"
	"runtime"
	"time"
)

type pool struct {
	WorkPool chan *Worker
	// Job             chan *Worker
	maxWorkNum      int
	maxWorkDuration time.Duration
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

	p.WorkPool = make(chan *Worker, p.maxWorkNum)

	p.dispatch()

	return p, nil
}

func (p *pool) Add(w *Worker) {
	w.p = p
	p.WorkPool <- w
}

func (p *pool) dispatch() {
	go func() {
		for {
			select {
			case w := <-p.WorkPool:
				w.do()
			default:
				w := <-p.WorkPool
				w.do()
			}
		}
	}()
}
