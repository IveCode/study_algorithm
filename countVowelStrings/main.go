package main

import "fmt"

func main() {
	//fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings2(33))
}

func countVowelStrings(n int) int {
	dp := make([][5]int, n+1)
	for i := range dp {
		dp[i][0] = 1
	}

	for i := range dp {
		for j := 1; j < 5; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	printDBTable(dp)

	return dp[n][4]
}

func printDBTable(dp [][5]int) {
	for i := range dp {
		fmt.Println(dp[i])
	}
}

func countVowelStrings2(n int) int {
	dp := [5]int{5, 4, 3, 2, 1}
	for i := 1; i < n; i++ {
		dp[4] = dp[4]
		dp[3] += dp[4]
		dp[2] += dp[3]
		dp[1] += dp[2]
		dp[0] += dp[1]
	}
	return dp[0]
}
