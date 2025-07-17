package main

import (
	"fmt"
	"runtime"
	"sync"
)

func CalculateSqrs(values []int) {
	generator := func(done <-chan struct{}, values []int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, value := range values {
				select {
				case intStream <- value:
				case <-done:
					return
				}
			}
		}()
		return intStream
	}
	squarer := func(done <-chan struct{}, valuesStream <-chan int) <-chan int {
		resultStream := make(chan int)
		go func() {
			defer close(resultStream)
			for value := range valuesStream {
				select {
				case resultStream <- value * value:
				case <-done:
					return
				}
			}
		}()
		return resultStream
	}

	fanIn := func(done <-chan struct{}, channels ...<-chan int) <-chan int {
		resultingStream := make(chan int)
		var wg sync.WaitGroup
		multiplex := func(channel <-chan int) {
			defer wg.Done()
			for value := range channel {
				select {
				case <-done:
					return
				case resultingStream <- value:
				}
			}
		}
		wg.Add(len(channels))
		for _, channel := range channels {
			go multiplex(channel)
		}
		go func() {
			wg.Wait()
			close(resultingStream)
		}()
		return resultingStream
	}

	done := make(chan struct{})
	defer close(done)
	intStream := generator(done, values)

	numWorkers := runtime.NumCPU()
	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = squarer(done, intStream)
	}

	pipeline := fanIn(done, workers...)
	for v := range pipeline {
		fmt.Println(v)
	}
}
