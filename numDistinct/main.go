package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	var num int
	if len(s) < len(t) {
		return num
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][m] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	printDP(dp)
	return dp[0][0]
}

func printDP(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
}
