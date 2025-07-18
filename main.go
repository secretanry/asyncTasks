package main

import "fmt"

func main() {
	intersection := Intersection([]int{1, 2, 3}, []int{2, 3, 4})
	for _, val := range intersection {
		fmt.Println(val)
	}
}
