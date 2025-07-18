package main

import (
	"reflect"
	"testing"
)

func TestTask11(t *testing.T) {
	tests := []struct {
		a, b   []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 1, 2, 3}, []int{1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{}},
		{[]int{1, 2, 2, 3, 4}, []int{2, 2, 4, 4}, []int{2, 2, 4}},
	}
	for _, tt := range tests {
		res := Intersection(tt.a, tt.b)
		if !reflect.DeepEqual(res, tt.expect) {
			t.Errorf("Intersection(%v, %v) = %v, want %v", tt.a, tt.b, res, tt.expect)
		}
	}
}
