package main

func BinarySearch(a []int, k int) int {
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == k {
			return mid
		} else if a[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
