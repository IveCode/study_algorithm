package main

import "fmt"

//https://leetcode-cn.com/problems/house-robber/submissions/
//198. 打家劫舍
/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
    偷窃到的最高金额 = 2 + 9 + 1 = 12 。


提示：

0 <= nums.length <= 100
0 <= nums[i] <= 400
*/
func main() {
	fmt.Println("rob1:")
	fmt.Println(rob1([]int{1, 2, 3, 1}))
	fmt.Println(rob1([]int{2, 7, 9, 3, 1}))

	fmt.Println("rob2:")
	fmt.Println(rob2([]int{1, 2, 3, 1}))
	fmt.Println(rob2([]int{2, 7, 9, 3, 1}))

	fmt.Println("rob3:")
	fmt.Println(rob3([]int{1, 2, 3, 1}))
	fmt.Println(rob3([]int{2, 7, 9, 3, 1}))
}

func rob3(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	v0, v1, v2 := 0, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		v0 = max(v1, nums[i]+v2)
		v2 = v1
		v1 = v0
	}
	return v0
}

func rob2(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}

func rob1(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp func([]int, int) int
	dp = func(nums []int, idx int) int {
		if len(nums) <= idx {
			return 0
		}

		value := max(dp(nums, idx+1), nums[idx]+dp(nums, idx+2))
		return value
	}

	return dp(nums, 0)
}
