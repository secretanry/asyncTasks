package main

import "testing"

func TestTask17(t *testing.T) {
	tests := []struct {
		a      []int
		k      int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
		{[]int{}, 1, -1},
	}
	for _, tt := range tests {
		res := BinarySearch(tt.a, tt.k)
		if res != tt.expect {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.a, tt.k, res, tt.expect)
		}
	}
}
