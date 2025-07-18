package main

func SetBitInt64(n int64, i uint, bit uint) int64 {
	i--
	if bit == 0 {
		return n &^ (1 << i)
	}
	return n | (1 << i)
}
