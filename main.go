package main

import "fmt"

func main() {
	a := 2
	b := 3
	fmt.Printf("Initial values: a=%d b=%d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("Swapped values: a=%d b=%d\n", a, b)
}
