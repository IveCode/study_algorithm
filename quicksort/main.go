package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 7, 4, 6, 9, 8, 5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, l, h int) {
	if l >= h {
		return
	}

	pivot := nums[l]
	i, j := l, h
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot

	quickSort(nums, l, i-1)
	quickSort(nums, i+1, h)
}
