package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber136([]int{2, 2, 1}))
	fmt.Println(singleNumber136([]int{4, 1, 2, 1, 2}))

	fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber137([]int{0, 1, 0, 1, 0, 1, 99}))

	fmt.Println(singleNumber260([]int{1, 2, 1, 3, 2, 5}))
}

//136. 只出现一次的数字  https://leetcode-cn.com/problems/single-number/
/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func singleNumber136(nums []int) int {
	var value int

	for _, num := range nums {
		value ^= num
	}

	return value
}

/*
137. 只出现一次的数字 II
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99
*/

func singleNumber137(nums []int) int {
	var once, twice int
	for _, num := range nums {
		once = (^twice) & (once ^ num)
		twice = (^once) & (twice ^ num)
	}

	return once
}

/*
260. 只出现一次的数字 III
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :

输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
*/

func singleNumber260(nums []int) []int {
	values := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := values[num]; ok {
			delete(values, num)
			continue
		}
		values[num] = struct{}{}
	}
	var result []int
	for k, _ := range values {
		result = append(result, k)
	}
	if len(result) == 2 {
		return result
	}
	return nil
}
