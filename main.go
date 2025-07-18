package main

import (
	"fmt"
)

func main() {
	result := MakeUnique([]string{"cat", "cat", "dog", "cat", "tree"})
	for _, s := range result {
		fmt.Println(s)
	}
}
