package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func TimerQuit(seconds int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds))
	defer cancel()
	dataStream := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Writer exited")
				return
			case dataStream <- rand.Int():
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Reader exited")
				return
			case val := <-dataStream:
				fmt.Printf("Reader received: %d\n", val)
			}
		}
	}()
	wg.Wait()
}
