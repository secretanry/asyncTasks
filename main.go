package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "12345678901234567890"
	b := "98765432109876543210"

	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)

	fmt.Println("BigAddInt:", BigAddInt(x, y))
	fmt.Println("BigSubInt:", BigSubInt(y, x))
	fmt.Println("BigMulInt:", BigMulInt(x, y))
	fmt.Println("BigDivInt:", BigDivInt(y, x))
}
