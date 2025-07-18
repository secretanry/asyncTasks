package main

import (
	"testing"
	"time"
)

func TestTask25(t *testing.T) {
	start := time.Now()
	Sleep(200 * time.Millisecond)
	elapsed := time.Since(start)
	if elapsed < 180*time.Millisecond || elapsed > 500*time.Millisecond {
		t.Errorf("Sleep(200ms) elapsed = %v, want ~200ms", elapsed)
	}
}
