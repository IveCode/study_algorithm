package main

import "fmt"

func main() {
	fmt.Println(message([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}

func message(nums []int) int {
	a, b := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		c := b
		if c < a+nums[i] {
			c = a + nums[i]
		}
		a, b = b, c
		fmt.Println("AB:", i, a, b)
	}

	if a > b {
		return a
	}
	return b
}

func message2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]

	for i := 1; i < length; i++ {
		tdp0 := dp0
		if tdp0 < dp1 { // 计算 dp[i][0]
			tdp0 = dp1
		}
		tdp1 := dp0 + nums[i] // 计算 dp[i][1]

		dp0 = tdp0 // 用 dp[i][0] 更新 dp_0
		dp1 = tdp1 // 用 dp[i][1] 更新 dp_1

	}
	if dp0 > dp1 {
		return dp0
	}
	return dp1
}
