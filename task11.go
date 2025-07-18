package main

import "sort"

func Intersection(a []int, b []int) []int {
	result := make([]int, 0)
	sort.Ints(a)
	sort.Ints(b)
	i1 := 0
	i2 := 0
	for i1 < len(a) && i2 < len(b) {
		if a[i1] == b[i2] {
			result = append(result, a[i1])
			i1++
			i2++
			continue
		}
		if a[i1] > b[i2] {
			i2++
		} else {
			i1++
		}
	}
	return result
}
