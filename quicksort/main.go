package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 3, 2, 7, 4, 6, 9, 8, 5}
	nums := []int{2, 5, 4, 3, 2, 1}
	QuickSort(nums)
	fmt.Println(nums)
}

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	pivotIdx := 0
	i, j := 1, len(nums)-1
	for pivotIdx < j {
		if nums[i] <= pivot {
			nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]
			pivotIdx++
			i++
		} else {
			nums[j], nums[i] = nums[i], nums[j]
			j--
		}
	}
	QuickSort(nums[:pivotIdx])
	QuickSort(nums[pivotIdx+1:])
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
