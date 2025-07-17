package main

import (
	"os"
	"os/signal"
	"syscall"
)

func SigintQuit() {
	notifyChan := make(chan os.Signal, 1)
	signal.Notify(notifyChan, syscall.SIGINT)
	defer close(notifyChan)
	defer signal.Stop(notifyChan)
	// Some infinite work
	go func() {
		for {
			select {
			case <-notifyChan:
			default:
				// Meaningful work
			}
		}
	}()
	<-notifyChan
}
