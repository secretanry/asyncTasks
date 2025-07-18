package main

import "testing"

func TestTask13(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("Swap failed: got a=%d, b=%d; want a=2, b=1", a, b)
	}

	x, y := -5, 100
	Swap(&x, &y)
	if x != 100 || y != -5 {
		t.Errorf("Swap failed: got x=%d, y=%d; want x=100, y=-5", x, y)
	}
}
