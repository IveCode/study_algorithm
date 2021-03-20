package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("aaaaaaaaa"))
	fmt.Println(removeDuplicates("aaaaaaaa"))
}

func removeDuplicates(S string) string {
	stack := []byte{S[0]}
	for i := 1; i < len(S); i++ {
		n := len(stack)
		if n > 0 && stack[n-1] == S[i] {
			stack = stack[0 : n-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
