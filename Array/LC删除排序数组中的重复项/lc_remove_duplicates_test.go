package lc__test

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	printRemoveDuplicates(t, []int{1, 1, 2})
	printRemoveDuplicates(t, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	printRemoveDuplicates(t, []int{1, 2, 2})
}

func printRemoveDuplicates(t *testing.T, nums []int) {
	t.Log("before:", nums)
	length := RemoveDuplicates(nums)
	t.Log("after:", nums, length)
}

/*
1. 前一个值与当前值比较，相同移动当前值，否则记录到j
2. 数量掌握线性表的下标处理方式
*/
func RemoveDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}
