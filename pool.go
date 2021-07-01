package hzypool

type pool struct {
	MaxWorkNum int
	WorkPool   chan *Worker
}

func New(sl ...Setter) *pool {
	p := &pool{}
	for _, s := range sl {
		s(p)
	}

	p.WorkPool = make(chan *Worker, p.MaxWorkNum)

	return p
}

type Setter func(p *pool)

func WithSetMaxNum(num int) Setter {
	return func(p *pool) {
		p.MaxWorkNum = num
	}
}

func (p *pool) Add(w *Worker) {
	p.WorkPool <- w
}

func (p *pool) dispatch() {
	go func() {
		select {
		case w := <-p.WorkPool:
			w.do()
		}
	}()
}

func (p *pool) Run() {
	p.dispatch()
}
