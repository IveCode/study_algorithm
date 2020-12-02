package lc_test

import (
	"sort"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	printFindKthLargest(t, []int{3, 2, 1, 5, 6, 4}, 2)
	printFindKthLargest(t, []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	printFindKthLargest(t, []int{3}, 1)
}

func printFindKthLargest(t *testing.T, nums []int, k int) {
	t.Log("before:", nums, k)
	value := FindKthLargest2(nums, k)
	t.Log("after:", nums, k, value)
}

/*
	1. 先排序，后索引
*/
func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func FindKthLargest2(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return 0
	}

	index := FindKthLargestValue(nums, 0, n-1, k)
	return nums[index]
}

func partition(nums []int, left, right int) int {
	v := nums[left]
	for left < right {
		for left < right && v >= nums[right] {
			right -= 1
		}
		if left < right {
			nums[left] = nums[right]
		}

		for left < right && v <= nums[left] {
			left += 1
		}
		if left < right {
			nums[right] = nums[left]
		}
	}
	nums[left] = v
	return left
}
func FindKthLargestValue(nums []int, left, right, k int) int {
	if left >= right {
		return left
	}
	index := partition(nums, left, right)
	if index+1 > k {
		return FindKthLargestValue(nums, left, index-1, k)
	}
	if index+1 < k {
		return FindKthLargestValue(nums, index+1, right, k)
	}
	return index
}
