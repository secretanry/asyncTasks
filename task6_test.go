package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestTask6(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log.SetOutput(os.Stdout)
	GoStop()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	checks := []string{
		"[chan] exited by done",
		"[ctxCancel] exited by cancel",
		"[ctxTimeout] exited by timeout",
		"[ctxDeadline] exited by deadline",
		"[timer] exited by timer",
		"[goexit] exited by Goexit",
		"All goroutinnes stopped",
	}
	for _, check := range checks {
		if !strings.Contains(output, check) {
			t.Errorf("Output does not contain: %s\nFull output:\n%s", check, output)
		}
	}
}
