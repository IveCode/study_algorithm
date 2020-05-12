package main

import "fmt"

//1248. 统计「优美子数组」 https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
func main() {
	fmt.Println(numberOfSubarrays2([]int{1, 1, 2, 1, 1}, 3))
	fmt.Println(numberOfSubarrays2([]int{2, 4, 6}, 1))
	fmt.Println(numberOfSubarrays2([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}

func numberOfSubarrays(nums []int, k int) int {
	dp := make([]int, 0)
	cnt, ret := 0, 0
	for i := 0; i < len(nums); i++ {
		cnt++
		if nums[i]%2 == 1 {
			dp = append(dp, cnt)
			cnt = 0
		}
		if len(dp) >= k {
			ret += dp[len(dp)-k]
			fmt.Println(ret, dp)
		}
	}
	return ret
}

func numberOfSubarrays2(nums []int, k int) int {
	var (
		result int
		odd    []int
		cnt    int
	)
	for _, num := range nums {
		cnt++
		if num&1 == 1 {
			odd = append(odd, cnt)
			cnt = 0
		}

		if len(odd) >= k {
			result += odd[len(odd)-k]
		}
	}

	return result
}
