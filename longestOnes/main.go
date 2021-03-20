package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode-cn.com/problems/max-consecutive-ones-iii/
1004. 最大连续1的个数 III
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

示例 1：

输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。


提示：
1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1

*/

func main() {
	fmt.Println(longestOnes2([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes2([]int{0, 0, 0, 0, 0}, 2))
	fmt.Println(longestOnes2([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(longestOnes2([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}

func longestOnes2(A []int, K int) int {
	P := make([]int, len(A)+1)
	for i, v := range A {
		P[i+1] = P[i] + 1 - v
	}
	var result int
	for right, v := range P {
		left := sort.SearchInts(P, v-K)
		result = max(result, right-left)
	}
	return result
}

func longestOnes(A []int, K int) int {
	var lsum, rsum, left, result int
	for rigth, v := range A {
		rsum += 1 - v
		for lsum < rsum-K {
			lsum += 1 - A[left]
			left++
		}

		result = max(result, rigth-left+1)
		fmt.Println(lsum, rsum, left, rigth, result)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
