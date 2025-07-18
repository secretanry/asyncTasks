package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

func GoStop() {
	var wg sync.WaitGroup

	// 1. Channel stop
	wg.Add(1)
	done := make(chan struct{})
	go func() {
		log.Println("[chan] started")
		defer wg.Done()
		for {
			select {
			case <-done:
				log.Println("[chan] exited by done")
				return
			default:
				log.Println("[chan] working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 2. ctx with cancel stop
	wg.Add(1)
	ctxCancel, cancelCancel := context.WithCancel(context.Background())
	go func() {
		log.Println("[ctxCancel] started")
		defer wg.Done()
		for {
			select {
			case <-ctxCancel.Done():
				log.Println("[ctxCancel] exited by cancel")
				return
			default:
				log.Println("[ctxCancel] working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 3. ctx with timeout stop
	wg.Add(1)
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelTimeout()
	go func() {
		log.Println("[ctxTimeout] started")
		defer wg.Done()
		for {
			select {
			case <-ctxTimeout.Done():
				log.Println("[ctxTimeout] exited by timeout")
				return
			default:
				log.Println("[ctxTimeout] working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 4. ctx with deadline stop
	wg.Add(1)
	ctxDeadline, deadlineCancel := context.WithDeadline(ctxTimeout, time.Now().Add(1*time.Second))
	defer deadlineCancel()
	go func() {
		log.Println("[ctxDeadline] started")
		defer wg.Done()
		for {
			select {
			case <-ctxDeadline.Done():
				log.Println("[ctxDeadline] exited by deadline")
				return
			default:
				log.Println("[ctxDeadline] working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 5. timer exit
	wg.Add(1)
	go func() {
		log.Println("[timer] started")
		defer wg.Done()
		timeout := time.After(1 * time.Second)
		for {
			select {
			case <-timeout:
				log.Println("[timer] exited by timer")
				return
			default:
				log.Println("[timer] working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 6. runtime.Goexit()
	wg.Add(1)
	go func() {
		log.Println("[goexit] started")
		defer wg.Done()
		log.Println("[goexit] working...")
		time.Sleep(1 * time.Second)
		log.Println("[goexit] exited by Goexit")
		runtime.Goexit()
	}()

	close(done)
	cancelCancel()
	wg.Wait()
	log.Println("All goroutinnes stopped")
}
