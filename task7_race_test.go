//go:build !race

package main

import "testing"

func TestTask7(t *testing.T) {
	t.Skip("TestTask7 is skipped under -race due to demonstration purposes")
}
