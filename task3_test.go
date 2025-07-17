package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTask3(t *testing.T) {
	done := make(chan struct{})
	inputStream := make(chan int)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	numWorkers := 3
	wp := NewWorkerPool(done, numWorkers, inputStream)
	wp.Start()

	values := []int{1, 2, 3, 4, 5}
	go func() {
		for _, v := range values {
			inputStream <- v
		}
		close(inputStream)
	}()

	time.Sleep(200 * time.Millisecond)
	close(done)
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	found := make(map[int]bool)
	for _, line := range strings.Split(output, "\n") {
		if v, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			found[v] = true
		}
	}

	for _, v := range values {
		if !found[v] {
			t.Errorf("Output does not contain value: %d", v)
		}
	}
}
