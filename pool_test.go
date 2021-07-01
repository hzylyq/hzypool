package hzypool_test

import (
	"context"
	"github.com/hzylyq/hzypool"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	p := hzypool.New(hzypool.WithSetMaxNum(10))

	p.Add(&hzypool.Worker{
		Fn:  fn,
		Arg: nil,
	})
	p.Run()
}

func fn(ctx context.Context, arg interface{}) error {
	time.Sleep(1 * time.Second)
	return nil
}
