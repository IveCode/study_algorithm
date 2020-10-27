package lc__test

import "testing"

func TestSetColors(t *testing.T) {
	printSetColors(t, []int{2, 0, 2, 1, 1, 0})
	printSetColors(t, []int{2, 1, 0, 2, 1, 0, 1, 0})
	printSetColors(t, []int{0, 0, 2, 1, 0, 2, 1, 0, 1, 0, 2, 2})
	printSetColors(t, []int{0})
	printSetColors(t, []int{0, 0})
	printSetColors(t, []int{1, 1})
	printSetColors(t, []int{2, 2})
	printSetColors(t, []int{2, 1})
	printSetColors(t, []int{1, 2, 0})
}

func printSetColors(t *testing.T, nums []int) {
	t.Log("before:", nums)
	setColors(nums)
	t.Log("after:", nums)
}

/*
基数排序法
可以采用基数排序法的思想，用一个数组记录下 0，1，3 的次数，后重排，这个算法对数组进行了两次遍历，其实有一种只进行一次遍历的方法。

三路快速排序方法
设置三个 lt, gt, i 定义：nums[0...lt] == 0，nums[lt+1...i-1] = 1，nums[gt...n-1] == 2，遍历一遍改数列保持这个定义。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/all-about-array/x9dam3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func setColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	l, r := 0, len(nums)-1
	for i := 0; i <= r; {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}
	}
}
