package main

import "fmt"

//22. 括号生成 https://leetcode-cn.com/problems/generate-parentheses/
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	p := NewMatchParenthesis()
	p.dfs(0, n, n, n)
	return p.results
}

type MatchParenthesis struct {
	results     []string
	parenthesis string
}

func NewMatchParenthesis() *MatchParenthesis {
	return &MatchParenthesis{results: []string{}}
}

func (p *MatchParenthesis) dfs(i, left, right, n int) {
	if i == 2*n {
		p.results = append(p.results, p.parenthesis)
		return
	}

	if left > 0 {
		p.parenthesis += "("
		p.dfs(i+1, left-1, right, n)
		p.parenthesis = p.parenthesis[:len(p.parenthesis)-1]
	}

	if right > 0 && right > left {
		p.parenthesis += ")"
		p.dfs(i+1, left, right-1, n)
		p.parenthesis = p.parenthesis[:len(p.parenthesis)-1]
	}
}
