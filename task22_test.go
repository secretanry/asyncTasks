package main

import (
	"math/big"
	"testing"
)

func TestTask22(t *testing.T) {
	a := "12345678901234567890"
	b := "98765432109876543210"

	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)

	if BigAddInt(x, y).String() != "111111111011111111100" {
		t.Errorf("BigAddInt failed")
	}
	if BigSubInt(y, x).String() != "86419753208641975320" {
		t.Errorf("BigSubInt failed")
	}
	if BigMulInt(x, y).String() != "1219326311370217952237463801111263526900" {
		t.Errorf("BigMulInt failed")
	}
	if BigDivInt(y, x).Cmp(big.NewInt(8)) != 0 {
		t.Errorf("BigDivInt failed: got %s, want 8", BigDivInt(y, x).String())
	}
}
