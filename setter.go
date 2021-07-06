package hzypool

import "time"

type Setter func(p *pool)

func WithSetMaxNum(num int) Setter {
	return func(p *pool) {
		p.maxWorkNum = num
	}
}

func WithSetMaxWorkDuration(duration time.Duration) Setter {
	return func(p *pool) {
		p.maxWorkDuration = duration
	}
}
