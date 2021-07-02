package hzypool

import (
	"fmt"
	"log"
	"runtime"
)

type pool struct {
	maxWorkNum int
	WorkPool   chan *Worker
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

	return p, nil
}

type Setter func(p *pool)

func WithSetMaxNum(num int) Setter {
	return func(p *pool) {
		p.maxWorkNum = num
	}
}

func (p *pool) Add(w *Worker) {
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
