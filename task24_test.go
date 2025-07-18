package main

import (
	"math"
	"testing"
)

func TestTask24(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	if d := p1.Distance(p2); math.Abs(d-5) > 1e-9 {
		t.Errorf("Distance = %v, want 5", d)
	}

	p3 := NewPoint(-1, -1)
	p4 := NewPoint(2, 3)
	if d := p3.Distance(p4); math.Abs(d-5) > 1e-9 {
		t.Errorf("Distance = %v, want 5", d)
	}

	p5 := NewPoint(1.5, 2.5)
	p6 := NewPoint(1.5, 2.5)
	if d := p5.Distance(p6); d != 0 {
		t.Errorf("Distance = %v, want 0", d)
	}
}
