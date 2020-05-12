package main

import "fmt"

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
