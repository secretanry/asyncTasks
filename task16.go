package main

func QuickSort(a []int) []int {
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		quickSort(a, low, pivot-1)
		quickSort(a, pivot+1, high)
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
