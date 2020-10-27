package lc__test

import "testing"

func TestRemoveElement(t *testing.T) {
	printRemoveElement(t, []int{3, 2, 2, 3}, 3)
	printRemoveElement(t, []int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}

func printRemoveElement(t *testing.T, nums []int, val int) {
	t.Log("after:", val, nums)
	length := RemoveElement(nums, val)
	t.Log("before:", val, nums, length)
}

func RemoveElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}

	return j
}
