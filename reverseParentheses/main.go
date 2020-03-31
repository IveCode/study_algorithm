package main

import "fmt"

func main() {
	fmt.Println(reverseParentheses("(abcd)"))
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}

func reverseParentheses(s string) string {
	var leftBrackets []int
	ss := []byte(s)
	for i := 0; i < len(s); i++ {
		if ss[i] == '(' {
			leftBrackets = append(leftBrackets, i) //记录位置
		}

		if ss[i] == ')' {
			j := leftBrackets[len(leftBrackets)-1] + 1
			k := i - 1
			for j < k {
				ss[j], ss[k] = ss[k], ss[j]
				j++
				k--
			}
			leftBrackets = leftBrackets[:len(leftBrackets)-1]
		}
	}

	var result string
	for _, v := range ss {
		if v == '(' || v == ')' {
			continue
		}
		result += string(v)
	}

	return result
}
