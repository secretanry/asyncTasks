//go:build race

package main

import "testing"

func TestTask4(t *testing.T) {
	t.Skip("TestTask4 is skipped under -race due to subprocess limitations")
}
