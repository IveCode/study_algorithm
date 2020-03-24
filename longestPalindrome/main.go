package main

import "fmt"

//最长回文子串 leecode: https://leetcode-cn.com/problems/longest-palindromic-substring/
func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		length := palindromeStringLength(s, i, i)
		length2 := palindromeStringLength(s, i, i+1)
		if length < length2 {
			length = length2
		}
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func palindromeStringLength(s string, l, r int) int {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return r - l - 1
}
