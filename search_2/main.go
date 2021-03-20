package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	mid := (right-left)/2 + left

	if nums[mid] == target {
		return mid
	}

	if nums[left] < nums[mid] && nums[left] <= target && target <= nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	}

	if nums[left] > nums[mid] && (nums[left] <= target || target < nums[mid]) {
		return binarySearch(nums, left, mid-1, target)
	}

	if nums[mid] < nums[right] && nums[mid] < target && target <= nums[right] {
		return binarySearch(nums, mid+1, right, target)
	}

	if nums[mid] > nums[right] && (nums[mid] < target || target <= nums[right]) {
		return binarySearch(nums, mid+1, right, target)
	}
	return -1
}
