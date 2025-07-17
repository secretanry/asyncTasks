package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTask2(t *testing.T) {
	values := []int{2, 4, 6, 8, 10}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CalculateSqrs(values)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := []string{"4", "16", "36", "64", "100"}
	for _, sqr := range expected {
		if !strings.Contains(output, sqr+"\n") && !strings.Contains(output, sqr+"\r\n") {
			t.Errorf("Output does not contain square: %s", sqr)
		}
	}
}
