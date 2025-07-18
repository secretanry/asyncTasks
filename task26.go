package main

import "unicode"

func UniqueChars(s string) bool {
	seen := make(map[rune]struct{})
	for _, r := range s {
		lower := unicode.ToLower(r)
		if _, ok := seen[lower]; ok {
			return false
		}
		seen[lower] = struct{}{}
	}
	return true
}
