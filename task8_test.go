package main

import "testing"

func TestTask8(t *testing.T) {
	tests := []struct {
		n   int64
		i   uint
		bit uint
		exp int64
	}{
		{5, 1, 0, 4},  // 0101 -> 0100 (сбросить 1-й бит справа)
		{5, 2, 0, 5},  // 0101 -> 0101 (2-й бит уже 0)
		{5, 3, 0, 1},  // 0101 -> 0001 (сбросить 3-й бит)
		{5, 3, 1, 5},  // 0101 -> 0101 (уже 1)
		{4, 1, 1, 5},  // 0100 -> 0101 (установить 1-й бит)
		{0, 4, 1, 8},  // 0000 -> 1000 (установить 4-й бит)
		{15, 4, 0, 7}, // 1111 -> 0111 (сбросить 4-й бит)
	}
	for _, tt := range tests {
		res := SetBitInt64(tt.n, tt.i, tt.bit)
		if res != tt.exp {
			t.Errorf("SetBitInt64(%d, %d, %d) = %d, want %d", tt.n, tt.i, tt.bit, res, tt.exp)
		}
	}
}
