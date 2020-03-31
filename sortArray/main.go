package main

import "fmt"

//912. 排序数组 https://leetcode-cn.com/problems/sort-an-array/

func main() {
	fmt.Println(sortArray([]int{5, 4, 3, 2, 1}))
}

func sortArray(nums []int) []int {

	//return InsertSort(nums)
	quickSort(nums)
	return nums
}

func merge(nums, left, right []int) {
	i, l, r := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			nums[i] = left[l]
			l++
		} else {
			nums[i] = right[r]
			r++
		}
		i++
	}

	for l < len(left) {
		nums[i] = left[l]
		l++
		i++
	}

	for r < len(right) {
		nums[i] = right[l]
		l++
		r++
	}
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func InsertSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		e := nums[i]
		j := i
		for ; j > 0 && e < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
	return nums
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	reference := nums[0]
	head, tail := 0, len(nums)-1
	i := 1
	for head < tail {
		if nums[i] <= reference {
			nums[head], nums[i] = nums[i], nums[head]
			head++
			i++
		} else {
			nums[tail], nums[i] = nums[i], nums[tail]
			tail--
		}
	}
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}
