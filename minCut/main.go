package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("result:", minCut("abc"))
}

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	printDPTable("init", dp)
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1])
		}
	}
	printDPTable("proc", dp)

	results := make([]int, n)
	for i := 0; i < n; i++ {
		if dp[0][i] {
			results[i] = 0
			continue
		}
		results[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if dp[j+1][i] {
				results[i] = int(math.Min(float64(results[j]+1), float64(results[i])))
			}
		}
	}
	fmt.Println(results)
	return results[n-1]
}

func printDPTable(desc string, dp [][]bool) {
	fmt.Println(desc)
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
}
