package main

import (
	"testing"
)

func TestTask12(t *testing.T) {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	got := MakeUnique(input)
	expect := []string{"cat", "dog", "tree"}

	if !stringSlicesEqualUnordered(got, expect) {
		t.Errorf("MakeUnique(%v) = %v, want %v", input, got, expect)
	}
}

func stringSlicesEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for j, y := range b {
			if !used[j] && x == y {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
