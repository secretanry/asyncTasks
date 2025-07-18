package main

import (
	"reflect"
	"testing"
)

func TestTask16(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
	}
	for _, tt := range tests {
		in := make([]int, len(tt.input))
		copy(in, tt.input)
		res := QuickSort(in)
		if !reflect.DeepEqual(res, tt.expect) {
			t.Errorf("QuickSort(%v) = %v, want %v", tt.input, res, tt.expect)
		}
	}
}
