package main

import (
	"context"
	"sync"
	"time"
)

func ConcurrentMap() {
	concMap := sync.Map{}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				concMap.Store("key", "value1")
				concMap.Load("key")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				concMap.Store("key", "value2")
				concMap.Load("key")
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
}
