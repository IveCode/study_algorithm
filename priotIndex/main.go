package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))      //3
	fmt.Println(pivotIndex([]int{1, 2, 3}))               //-1
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0})) //2
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 0, 1}))  //1
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	lResult := 0
	for i, num := range nums {
		lResult += num
		rResult := 0
		for j := i; j < len(nums); j++ {
			rResult += nums[j]
		}
		if lResult == rResult {
			return i
		}
	}

	return -1
}
