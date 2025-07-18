package main

import (
	"fmt"
	"sort"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := GroupTemperatures(temps)

	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: ", k)
		for i, v := range groups[k] {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", v)
		}
		fmt.Println()
	}
}
