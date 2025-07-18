package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPipeline(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DoublingPipeline(values)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := []string{"2", "4", "6", "8", "10"}
	for _, want := range expected {
		if !strings.Contains(output, want+"\n") && !strings.Contains(output, want+"\r\n") {
			t.Errorf("Output does not contain: %s", want)
		}
	}
}
