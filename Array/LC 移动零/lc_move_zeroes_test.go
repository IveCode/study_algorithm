package lc_test

import "testing"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/all-about-array/x9rh8e/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func TestMoveZeroes(t *testing.T) {
	printMoveZeros(t, []int{0, 1, 0, 3, 12})
	printMoveZeros(t, []int{0, 0, 0, 0, 0})
	printMoveZeros(t, []int{2, 1})
	printMoveZeros(t, []int{0})
	printMoveZeros(t, []int{1})
}

func printMoveZeros(t *testing.T, nums []int) {
	t.Log("before:", nums)
	moveZeroes(nums)
	t.Log("after:", nums)
}

/*
1. 将数组中不等于0的数往前移;
2. j记录数组中不为0值的下标，i表示已经遍历的下标；
*/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}

func moveZeroes2(nums []int) {
	var zeroNums int
	for i := 0; i < len(nums)-zeroNums; {
		if nums[i] != 0 {
			i++
			continue
		}
		for j := i + 1; j < len(nums)-zeroNums; j++ {
			nums[j-1] = nums[j]
		}
		nums[len(nums)-zeroNums-1] = 0
		zeroNums += 1
	}
}
