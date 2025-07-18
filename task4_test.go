//go:build !race

package main

import (
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestTask4(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcessTask4")
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	time.Sleep(100 * time.Millisecond)
	if err := cmd.Process.Signal(syscall.SIGINT); err != nil {
		t.Fatalf("Failed to send SIGINT: %v", err)
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Process exited with error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Error("Process did not exit on SIGINT")
	}
}

func TestHelperProcessTask4(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	SigintQuit()
	os.Exit(0)
}
