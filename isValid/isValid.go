package main

import "fmt"

//20. 有效的括号 https://leetcode-cn.com/problems/valid-parentheses/
func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := NewStack()
	for i, v := range s {
		if i == 0 && (v == ')' || v == ']' || v == '}') {
			return false
		}

		if v == '(' || v == '[' || v == '{' {
			stack.Push(byte(v))
			continue
		}

		if v == ')' && stack.Front() == '(' {
			stack.Pop()
		} else if v == ']' && stack.Front() == '[' {
			stack.Pop()
		} else if v == '}' && stack.Front() == '{' {
			stack.Pop()
		} else {
			return false
		}
	}

	if stack.length != 0 {
		return false
	}

	return true
}

type Stack struct {
	value  []byte
	length int
}

func NewStack() *Stack {
	return &Stack{value: []byte{}}
}

func (s *Stack) Push(b byte) {
	s.value = append(s.value, b)
	s.length++
}

func (s *Stack) Pop() byte {
	s.length--
	v := s.value[s.length]
	s.value = s.value[:s.length]
	return v
}

func (s *Stack) Front() byte {
	return s.value[s.length-1]
}
