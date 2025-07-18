package main

func RemoveAt[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}
