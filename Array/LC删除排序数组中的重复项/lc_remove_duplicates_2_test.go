package lc__test

import "testing"

func TestRemoveDuplicatesII(t *testing.T) {
	printRemoveDuplicatesII(t, []int{1, 1, 1, 2, 2, 3})
	printRemoveDuplicatesII(t, []int{0, 0, 1, 1, 1, 1, 2, 3, 3})
}

func printRemoveDuplicatesII(t *testing.T, nums []int) {
	t.Log("before:", nums)
	length := RemoveDuplicatesII2(nums)
	t.Log("after:", nums, length)
}

//利用3个数比较进行窗口移动
//k,i 最近两个窗口的值，j表示当前的值
func RemoveDuplicatesII2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	i := 1
	k := i - 1
	j := i + 1
	for ; j <= n-1; j++ {
		if (nums[j] != nums[i]) || (nums[j] == nums[i] && nums[j] != nums[k]) {
			k = i
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

//详情：
// 1 1 1 1 2 3 3
// k i j

// 1 1 1 1 2 3 3
// k i   j

// 1 1 1 1 2 3 3
// k i     j

// 1 1 2 1 2 3 3
//   k i     j

// 1 1 2 3 2 3 3
//     k i     j

// 1 1 2 3 3 3 3 EOF
//       k i     j

//统计个数，满足条件的记录到下标index中
func RemoveDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count := 1
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
