package hzypool

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type pool struct {
	WorkPool        chan *Worker
	Job             chan *Worker
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
	p.Job = make(chan *Worker, 1)

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
				log.Print("case")
				w.do()
			default:
				log.Printf("default")
			}
		}
	}()
}

func (p *pool) Run() {
	p.dispatch()
}
