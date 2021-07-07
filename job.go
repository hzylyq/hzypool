package hzypool

import "context"

type fn func(ctx context.Context, arg interface{}) error

type Job struct {
	Fn  fn
	Arg interface{}
}

func NewJob(arg interface{}, fn fn) *Job {
	return &Job{
		Arg: arg,
		Fn:  fn,
	}
}
