package main

func ReverseWordsInPlace(s string) string {
	b := []byte(s)
	reverseBytes(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverseBytes(b, start, i-1)
			start = i + 1
		}
	}
	return string(b)
}

func reverseBytes(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
