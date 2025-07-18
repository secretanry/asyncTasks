package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	num atomic.Int32
}

func (c *AtomicCounter) Inc() {
	c.num.Add(1)
}

func (c *AtomicCounter) Load() int {
	return int(c.num.Load())
}

func ConcurrentIncrement() {
	done := make(chan struct{})
	defer close(done)
	var wg sync.WaitGroup
	counter := &AtomicCounter{}
	startIncrement := func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case <-done:
				return
			default:
				counter.Inc()
			}
		}
	}
	wg.Add(3)
	go startIncrement()
	go startIncrement()
	go startIncrement()
	wg.Wait()
	fmt.Println(counter.Load())
}
