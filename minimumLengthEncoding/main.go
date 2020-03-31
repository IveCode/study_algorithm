package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
	fmt.Println(minimumLengthEncoding([]string{"me", "time"}))
}

func minimumLengthEncoding1(words []string) int {
	set := map[string]struct{}{}
	for _, v := range words {
		set[v] = struct{}{}
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			if _, ok := set[word[i:]]; ok {
				delete(set, word[i:])
			}
		}
	}
	result := 0
	for k, _ := range set {
		result += len(k) + 1
	}
	return result
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) > len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func minimumLengthEncoding(words []string) int {
	sort.Sort(StringSlice(words))
	fmt.Println(words)
	tire := NewTire()
	result := 0
	for _, word := range words {
		result += tire.Insert(word)
	}
	return result
}

type TireNode struct {
	children [26]*TireNode
}

func NewTireNode() *TireNode {
	return &TireNode{}
}

type Tire struct {
	root *TireNode
}

func NewTire() *Tire {
	return &Tire{root: NewTireNode()}
}

func (t *Tire) Insert(word string) int {
	cur := t.root
	isNew := false
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i] - 'a'
		for cur.children[c] == nil {
			isNew = true
			cur.children[c] = NewTireNode()
		}
		cur = cur.children[c]
	}
	if isNew {
		return len(word) + 1
	}
	return 0
}
