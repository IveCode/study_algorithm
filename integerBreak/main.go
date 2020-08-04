package main

import (
	"fmt"
)

//URL: https://leetcode-cn.com/problems/integer-break/

func main() {
	//fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
		//fmt.Println(dp)
	}
	//fmt.Println(dp)
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
