package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestTask18(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ConcurrentIncrement()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	output = strings.TrimSpace(output)
	val, err := strconv.Atoi(output)
	if err != nil {
		t.Fatalf("Output is not an integer: %q", output)
	}
	if val != 300 {
		t.Errorf("Counter value = %d, want 300", val)
	}
}
