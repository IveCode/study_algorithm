package main

import "fmt"

func main() {
	//fmt.Println(validPalindrome("aba"))
	//fmt.Println(validPalindrome("abca"))
	//fmt.Println(validPalindrome("abcda"))
	//fmt.Println(validPalindrome("yd"))
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] == s[right] {
			left++
			right--
			continue
		} else {
			isLeftDelete, isRightDelete := true, true
			for i, j := left+1, right; i <= j; {
				if s[i] != s[j] {
					isLeftDelete = false
					break
				}
				i++
				j--
			}

			for i, j := left, right-1; i <= j; {
				if s[i] != s[j] {
					isRightDelete = false
					break
				}
				i++
				j--
			}

			return isLeftDelete || isRightDelete
		}
	}

	return true
}
