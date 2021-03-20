package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 7}, {6, 4}, {2, 3}}))
	fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}))
	//fmt.Println(maxEnvelopes([][]int{{30, 50}, {12, 2}, {3, 4}, {12, 15}}))
	//fmt.Println(maxEnvelopes([][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}))
	fmt.Println(maxEnvelopes([][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}))
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1])
	})
	fmt.Println(envelopes)
	dp := make([]int, len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(envelopes); i++ {
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	result := dp[0]
	for i := 1; i < len(dp); i++ {
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
