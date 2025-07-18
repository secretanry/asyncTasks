package main

import "fmt"

func main() {
	ints := []int{10, 20, 30, 40, 50}
	fmt.Println("Before RemoveAt:", ints)
	ints = RemoveAt(ints, 2)
	fmt.Println("After RemoveAt (index 2):", ints)

	words := []string{"cat", "dog", "fish", "tree"}
	fmt.Println("Before RemoveAt:", words)
	words = RemoveAt(words, 1)
	fmt.Println("After RemoveAt (index 1):", words)
}
