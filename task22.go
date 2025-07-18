package main

import (
	"math/big"
)

func BigAddInt(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func BigSubInt(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func BigMulInt(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func BigDivInt(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
