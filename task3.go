package main

import (
	"fmt"
)

type WorkerPool struct {
	numWorkers int
	start      chan struct{}
}

func NewWorkerPool(done <-chan struct{}, numWorkers int, inputStream <-chan int) *WorkerPool {
	startChan := make(chan struct{})
	for i := 0; i < numWorkers; i++ {
		go func() {
			<-startChan
			for {
				select {
				case <-done:
					return
				case v, ok := <-inputStream:
					if !ok {
						return
					}
					fmt.Println(v)
				}
			}
		}()
	}
	return &WorkerPool{
		numWorkers: numWorkers,
		start:      startChan,
	}
}

func (p *WorkerPool) Start() {
	close(p.start)
}

// snapshot of main to run this task
//func main() {
//	numWorkers := runtime.NumCPU()
//	if len(os.Args) > 1 {
//		if n, err := strconv.Atoi(os.Args[1]); err == nil && n > 0 {
//			numWorkers = n
//		}
//	}
//
//	done := make(chan struct{})
//	signalChan := make(chan os.Signal, 1)
//	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
//	defer close(done)
//	go func() {
//		select {
//		case <-done:
//			return
//		case <-signalChan:
//			done <- struct{}{}
//			close(signalChan)
//			return
//		}
//	}()
//	inputStream := make(chan int)
//	wp := NewWorkerPool(done, numWorkers, inputStream)
//	wp.Start()
//	for {
//		select {
//		case <-done:
//			return
//		case inputStream <- rand.Int():
//		}
//	}
//}
