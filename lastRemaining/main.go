package main

//面试题62. 圆圈中最后剩下的数字 https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(10, 17))
	fmt.Println(lastRemaining2(5, 3))
	fmt.Println(lastRemaining2(10, 17))
}

func lastRemaining(n, m int) int {
	return f(n, m)
}

func f(n, m int) int {
	if n == 1 {
		return 0
	}
	x := f(n-1, m)
	v := (m + x) % n
	fmt.Println(v, "= (", m, "+", x, ") %", n)
	return v
}

func lastRemaining2(n, m int) int {
	f := 0
	for i := 1; i != n+1; i++ {
		f = (m + f) % i
	}
	return f
}
