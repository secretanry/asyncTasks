package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTask21(t *testing.T) {
	legacy := &LegacyPrinter{}
	adapter := &LegacyPrinterAdapter{OldPrinter: legacy}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	res := adapter.Print("hello")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Legacy: hello") {
		t.Errorf("Output does not contain expected legacy print: %q", output)
	}
	if res != "Legacy: hello" {
		t.Errorf("Adapter returned %q, want %q", res, "Legacy: hello")
	}
}
