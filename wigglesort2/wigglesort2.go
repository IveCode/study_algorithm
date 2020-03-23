package main

import (
	"fmt"
	"sort"
)

func WiggleSort2(nums []int) {
	sort.Ints(nums)
	l := len(nums)
	j, k := l, (l+1)>>1
	t := make([]int, l)
	copy(t, nums)
	for i := 0; i < l; i++ {
		if i&1 == 1 {
			j--
			nums[i] = t[j]
		} else {
			k--
			nums[i] = t[k]
		}
	}
}

func WiggleSort3(nums []int) {
	sort.Ints(nums)

	var i, j int
	k := len(nums) - 1
	mid := k / 2
	for j < k {
		if nums[j] > nums[mid] {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else if nums[j] < nums[mid] {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		} else {
			j++
		}
	}

	var pos1, pos2 int
	if len(nums)%2 == 0 {
		pos1 = len(nums)/2 - 1
	} else {
		pos1 = len(nums) / 2
	}
	pos2 = len(nums) - 1
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			tmp[i] = nums[pos1]
			pos1--
		} else {
			tmp[i] = nums[pos2]
			pos2--
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}

func main() {
	numss := [][]int{
		[]int{1, 5, 1, 1, 6, 4},
		[]int{1, 3, 2, 2, 3, 1},
		[]int{4, 1, 3, 5},
		[]int{4, 1, 3, 5, 6},
	}
	for _, nums := range numss {
		//fmt.Println(nums)
		WiggleSort2(nums)
		fmt.Println(nums)
	}
}

/*
[1 6 1 5 1 4]
[2 3 1 3 1 2]
[3 5 1 4]
[4 6 3 5 1]


[1 4 1 6 1 5]
[2 3 1 3 1 2]
[3 4 1 5]
[4 5 3 6 1]
*/
