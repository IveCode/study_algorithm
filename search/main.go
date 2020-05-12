package main

import "fmt"

//33. 搜索旋转排序数组  https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 3}, 0))
	fmt.Println(search([]int{1, 3}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
}

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	return recursionSearch(nums, 0, length-1, target)
}

func recursionSearch(nums []int, low, high int, target int) int {
	if low > high {
		return -1
	}

	lValue, hValue := nums[low], nums[high]
	if lValue == target {
		return low
	}

	if hValue == target {
		return high
	}

	mid := (high-low)/2 + low
	mValue := nums[mid]
	if mValue == target {
		return mid
	}

	if lValue < mValue {
		if lValue < target && target < mValue {
			return recursionSearch(nums, low+1, mid-1, target)
		}
		return recursionSearch(nums, mid+1, high-1, target)
	}

	if mValue < target && target < hValue {
		return recursionSearch(nums, mid+1, high-1, target)
	}
	return recursionSearch(nums, low+1, mid-1, target)
}
