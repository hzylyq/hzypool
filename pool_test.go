package hzypool_test

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	pool := sync.Pool{}
	pool.Get()

	pool.Put(1)
}
