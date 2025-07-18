package main

import "fmt"

func DoublingPipeline(values []int) {
	generator := func(done <-chan struct{}, values []int) <-chan int {
		outStream := make(chan int, 1)
		go func() {
			defer close(outStream)
			for _, value := range values {
				select {
				case <-done:
					return
				case outStream <- value:
				}
			}
		}()
		return outStream
	}
	doubler := func(done <-chan struct{}, inputStream <-chan int) <-chan int {
		outStream := make(chan int, 1)
		go func() {
			defer close(outStream)
			for value := range inputStream {
				select {
				case <-done:
					return
				case outStream <- value * 2:
				}
			}
		}()
		return outStream
	}

	done := make(chan struct{})
	defer close(done)
	pipeline := doubler(done, generator(done, values))
	for v := range pipeline {
		fmt.Println(v)
	}
}
