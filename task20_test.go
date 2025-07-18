package main

import "testing"

func TestTask20(t *testing.T) {
	tests := []struct {
		in, expect string
	}{
		{"snow dog sun", "sun dog snow"},
		{"hello", "hello"},
		{"", ""},
		{"a b c d", "d c b a"},
		{"глав рыба кот", "кот рыба глав"},
		{"😀 😃 😄 😁", "😁 😄 😃 😀"},
	}
	for _, tt := range tests {
		res := ReverseWordsInPlace(tt.in)
		if res != tt.expect {
			t.Errorf("ReverseWordsInPlace(%q) = %q, want %q", tt.in, res, tt.expect)
		}
	}
}
