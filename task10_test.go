package main

import (
	"testing"
)

func TestTask10(t *testing.T) {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	got := GroupTemperatures(temps)

	expect := map[int][]float64{
		-20: {-25.4, -27.0, -21.0},
		10:  {13.0, 19.0, 15.5},
		20:  {24.5},
		30:  {32.5},
	}

	for k, want := range expect {
		gotVals := got[k]
		if !float64SlicesEqualUnordered(gotVals, want) {
			t.Errorf("For group %d: got %v, want %v", k, gotVals, want)
		}
	}
}

func float64SlicesEqualUnordered(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for j, y := range b {
			if !used[j] && (x == y) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
