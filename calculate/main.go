package main

import "fmt"

func main() {
	fmt.Println(calculate("1 + 2"))               //3
	fmt.Println(calculate(" 2-1 + 2 "))           //3
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) //23
}

func calculate(s string) int {
	result := 0
	ops := []int{1}
	sign := 1
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result += sign * num
		}
	}
	return result
}
