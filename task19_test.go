package main

import "testing"

func TestTask19(t *testing.T) {
	tests := []struct {
		in, expect string
	}{
		{"главрыба", "абырвалг"},
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"😀😃😄😁", "😁😄😃😀"},
		{"абв😀гд", "дг😀вба"},
	}
	for _, tt := range tests {
		res := ReverseString(tt.in)
		if res != tt.expect {
			t.Errorf("ReverseString(%q) = %q, want %q", tt.in, res, tt.expect)
		}
	}
}
