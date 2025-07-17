package main

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

func TestTask5(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	TimerQuit(1)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	matched, _ := regexp.MatchString(`(?m)Reader received: \d+`, output)
	if !matched {
		t.Errorf("No values received by reader. Output: %s", output)
	}
	if !regexp.MustCompile(`Writer exited`).MatchString(output) {
		t.Errorf("Writer did not exit message. Output: %s", output)
	}
	if !regexp.MustCompile(`Reader exited`).MatchString(output) {
		t.Errorf("Reader did not exit message. Output: %s", output)
	}
}
