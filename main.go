package main

import (
	"fmt"
	"math/rand"
)

func main() {
	toSort := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	rand.Shuffle(len(toSort), func(i, j int) {
		toSort[i], toSort[j] = toSort[j], toSort[i]
	})
	res := QuickSort(toSort)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
