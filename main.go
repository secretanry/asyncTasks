package main

import (
	"fmt"
)

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	d := p1.Distance(p2)
	fmt.Printf("Расстояние между точками (0,0) и (3,4): %.2f\n", d)
}
