package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 1))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 0, 3, 0)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 0, 3, target)
}

func nSum(nums []int, start, n, target int) [][]int {
	if n < 2 || len(nums) < n {
		return nil
	}
	if n == 2 {
		return twoSum(nums, start, target)
	}

	length := len(nums) - 1
	var results [][]int
	for i := start; i < length; {
		nSumResults := nSum(nums, i+1, n-1, target-nums[i])

		for j, _ := range nSumResults {
			results = append(results, append([]int{nums[i]}, nSumResults[j]...))
		}
		for v := nums[i]; i < len(nums) && v == nums[i]; i++ {
		}
	}
	return results
}

func twoSum(nums []int, start, target int) [][]int {
	lo, hi := start, len(nums)-1
	var results [][]int
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum < target {
			for v := nums[lo]; lo < hi && v == nums[lo]; lo++ {
			}
		} else if sum > target {
			for v := nums[hi]; lo < hi && v == nums[hi]; hi-- {
			}
		} else {
			results = append(results, []int{nums[lo], nums[hi]})
			for v := nums[lo]; lo < hi && v == nums[lo]; lo++ {
			}
			for v := nums[hi]; lo < hi && v == nums[hi]; hi-- {
			}
		}
	}
	return results
}
