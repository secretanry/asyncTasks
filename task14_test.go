package main

import "testing"

func TestTask14(t *testing.T) {
	tests := []struct {
		val    interface{}
		expect string
	}{
		{42, "int"},
		{"hello", "string"},
		{true, "bool"},
		{make(chan int), "chan"},
		{make(chan string), "chan"},
		{make(chan struct{}), "chan"},
		{make(chan bool), "chan"},
		{make(chan float64), "chan"},
		{(chan int)(nil), "chan"},
		{3.14, "unknown"},
		{[]int{1, 2, 3}, "unknown"},
	}
	for _, tt := range tests {
		res := IdentifyType(tt.val)
		if res != tt.expect {
			t.Errorf("IdentifyType(%T) = %q, want %q", tt.val, res, tt.expect)
		}
	}
}
