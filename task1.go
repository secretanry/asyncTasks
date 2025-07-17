package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s\n", h.name)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La ", lyrics)
}

type Action struct {
	Human
}
