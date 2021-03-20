package main

import (
	"fmt"

	"github.com/go-ethereum/common/math"
)

/*
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println("===========")
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	result := make([][]int, len(prices))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(prices))
	}
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			value := prices[j] - prices[i]
			result[i][j] = value
			if max < value {
				max = value
			}
		}
	}
	//fmt.Println("=============")
	return max
}

func maxProfit2(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	for i := 0; i < len(prices); i++ {
		fmt.Println(i, prices[i], " => ", dp[i])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit3(prices []int) int {
	v1, v2 := 0, math.MinInt32

	for i := 0; i < len(prices); i++ {
		v1 = max(v1, v2+prices[i])
		v2 = max(v2, -prices[i])
		fmt.Println(i, prices[i], v1, v2)
	}
	return v1
}
