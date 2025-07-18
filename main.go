package main

import "fmt"

func main() {
	cases := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"abcABC",
		"😀😃😄😁",
		"😀😃😄😀",
	}
	for _, s := range cases {
		fmt.Printf("%q: %v\n", s, UniqueChars(s))
	}
}
