package main

import "testing"

func TestTask26(t *testing.T) {
	tests := []struct {
		in     string
		expect bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
		{"", true},
		{"A", true},
		{"abcABC", false},
		{"😀😃😄😁", true},
		{"😀😃😄😀", false},
	}
	for _, tt := range tests {
		res := UniqueChars(tt.in)
		if res != tt.expect {
			t.Errorf("UniqueChars(%q) = %v, want %v", tt.in, res, tt.expect)
		}
	}
}
