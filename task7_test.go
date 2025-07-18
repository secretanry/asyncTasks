//go:build race

package main

import "testing"

func TestTask7(t *testing.T) {
	ConcurrentMap()
}
