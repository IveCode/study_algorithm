package main

import "fmt"

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
[1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	x, y := len(grid), len(grid[0])

	dp := make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, y)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < x; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < y; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[x-1][y-1]
}

func printGrid(grid [][]int) {
	x, y := len(grid), len(grid[0])
	for i := 0; i < x; i++ {
		fmt.Print("[")
		for j := 0; j < y; j++ {
			fmt.Printf("%d", grid[i][j])
			if j+1 != y {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}
