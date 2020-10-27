package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("000111000111"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("000111"))
}

func countBinarySubstrings(s string) int {
	result := 0
	var counts []int
	for i := 0; i < len(s); {
		c := s[i]
		count := 0
		for i < len(s) && c == s[i] {
			i++
			count++
		}
		counts = append(counts, count)
	}
	for i := 1; i < len(counts); i++ {
		result += min(counts[i], counts[i-1])
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
