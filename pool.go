package hzypool

type pool struct {
	MaxWorkNum int
	WorkPool   chan *worker
}

func New(sl ...Setter) *pool {
	p := &pool{}
	for _, s := range sl {
		s(p)
	}

	p.WorkPool = make(chan *worker, p.MaxWorkNum)

	return &pool{}
}

type Setter func(p *pool)

func WithSetMaxNum(num int) Setter {
	return func(p *pool) {
		p.MaxWorkNum = num
	}
}

func (p *pool) add(w *worker) {
	p.WorkPool <- w
}

func (p *pool) dispatch() {
	go func() {
		select {
		case w := <-p.WorkPool:
			j := new(job)
			w.submit(j)
		}
	}()
}
