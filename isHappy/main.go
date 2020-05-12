package main

import "fmt"

//202. 快乐数 https://leetcode-cn.com/problems/happy-number/

func main() {
	printlnIsHappy(19)
}

func printlnIsHappy(n int) {
	fmt.Printf("%d isHappy: %t\n", n, isHappy(n))
}

func isHappy(n int) bool {
	for i := 0; i < 100; i++ {
		n = SquareSum(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func SquareSum(n int) int {
	total := 0
	for n > 0 {
		b := n % 10
		total += b * b
		n /= 10
	}
	return total
}
