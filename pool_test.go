package hzypool_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hzylyq/hzypool"
)

func TestNew(t *testing.T) {
	p, err := hzypool.New(
		hzypool.WithSetMaxNum(10),
		hzypool.WithSetMaxWorkDuration(10*time.Second))
	assert.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 10; i++ {
		p.Add(&hzypool.Worker{
			Fn:  fn,
			Arg: nil,
		})
	}
	p.Run()
	wg.Done()
	wg.Wait()
}

func fn(ctx context.Context, arg interface{}) error {
	time.Sleep(11 * time.Second)
	return nil
}
