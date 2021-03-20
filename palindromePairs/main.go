package main

import (
	"fmt"
)

func main() {
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
}

//字典树
type Node struct {
	ch   [26]int
	flag int
}

var tree []Node

func palindromePairs(words []string) [][]int {
	var results [][]int
	tree = []Node{Node{ch: [26]int{}, flag: -1}}
	for i, word := range words {
		insert(word, i)
	}

	for i, word := range words {
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findword(word, 0, j-1)
				if leftId != -1 && leftId != i {
					results = append(results, []int{i, leftId})
				}
			}

			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findword(word, j, m-1)
				if rightId != -1 && rightId != i {
					results = append(results, []int{rightId, i})
				}
			}
		}
	}
	return results
}

func insert(word string, flag int) {
	add := 0
	for i := 0; i < len(word); i++ {
		v := int(word[i] - 'a')
		if tree[add].ch[v] == 0 {
			tree = append(tree, Node{ch: [26]int{}, flag: -1})
			tree[add].ch[v] = len(tree) - 1
		}
		add = tree[add].ch[v]
	}
	tree[add].flag = flag
}

func findword(s string, left, right int) int {
	add := 0
	for i := right; i >= left; i-- {
		x := int(s[i] - 'a')
		if tree[add].ch[x] == 0 {
			return -1
		}
		add = tree[add].ch[x]
	}
	return tree[add].flag
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < right-left+1; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

//暴力穷举
/*
func palindromePairs(words []string) [][]int {
	var results [][]int

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if IsPalindrome(words[i] + words[j]) {
				results = append(results, []int{i, j})
			}
			if IsPalindrome(words[j] + words[i]) {
				results = append(results, []int{j, i})
			}
		}
	}

	return results
}
*/
func IsPalindrome(value string) bool {
	i, j := 0, len(value)-1
	for i <= j {
		if value[i] != value[j] {
			return false
		}
		i++
		j--
	}
	return true
}
