package main

import "testing"

func TestTask1(t *testing.T) {
	action := new(Action)
	action.name = "Mikhail"
	action.age = 20
	action.SayHi()
	action.Sing("La lala")
}
