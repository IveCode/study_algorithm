package main

import "fmt"

func main() {
	//fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsValue := make([]int, len(words))
	result := make([]int, len(queries))

	for i, query := range queries {
		queriesValue := minWord(query)
		value := 0
		for j, word := range words {
			if wordsValue[j] == 0 {
				wordsValue[j] = minWord(word)
			}
			if queriesValue < wordsValue[j] {
				value++
			}
		}
		result[i] = value
	}
	return result
}

func minWord(s string) int {
	word := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if word == s[i] {
			count++
		} else if word > s[i] {
			word = s[i]
			count = 1
		}
	}
	return count
}
