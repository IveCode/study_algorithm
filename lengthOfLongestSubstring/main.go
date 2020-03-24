package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	set := map[byte]int{}
	for i, idx := 0, 0; i < len(s); i++ {
		if v, ok := set[s[i]]; ok {
			if v > idx {
				idx = v
			}
		}
		if i-idx+1 > max {
			max = i - idx + 1
		}
		set[s[i]] = i + 1
	}
	return max
}

//超时
func lengthOfLongestSubstring1(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isRepeatString(s, i, j) {
				length := j - i
				if length > max {
					max = length
				}
			}
		}
	}

	return max
}

func isRepeatString(s string, start, end int) bool {
	set := map[byte]struct{}{}
	for i := start; i < end; i++ {
		if _, ok := set[s[i]]; ok {
			return false
		}
		set[s[i]] = struct{}{}
	}
	return true
}
