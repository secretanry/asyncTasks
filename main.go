package main

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

func main() {
	numWorkers := runtime.NumCPU()
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
			numWorkers = n
		}
	}

	done := make(chan struct{})
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	defer close(done)
	go func() {
		select {
		case <-done:
			return
		case <-signalChan:
			done <- struct{}{}
			close(signalChan)
			return
		}
	}()
	inputStream := make(chan int)
	wp := NewWorkerPool(done, numWorkers, inputStream)
	wp.Start()
	for {
		select {
		case <-done:
			return
		case inputStream <- rand.Int():
		}
	}
}
