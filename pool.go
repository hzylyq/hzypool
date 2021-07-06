package hzypool

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

type pool struct {
	workPool        chan *Worker
	jobs            chan *Job
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

	p.workPool = make(chan *Worker, p.maxWorkNum)
	p.jobs = make(chan *Job, 1)

	p.dispatch()

	return p, nil
}

func (p *pool) Add(w *Worker) {
	w.p = p
	p.workPool <- w
}
func (p *pool) Submit(j *Job) {
	if j == nil {
		return
	}

	p.jobs <- j
}

func (p *pool) dispatch() {
	go func() {
		for j := range p.jobs {
			log.Println(j)
			for {
				select {
				case w := <-p.workPool:
					log.Println("case")

					w.do()
				default:
					w := <-p.workPool
					log.Println("default")

					w.do()
				}
			}
		}
	}()
}
