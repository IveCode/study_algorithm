package main

import (
	"fmt"
)

func main() {
	//fmt.Println(findMedianSortedArrays([]int{}, []int{1, 3, 4, 5}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{}))
	//fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 5}, []int{3, 4, 6}))
	fmt.Println("===========")
	fmt.Println(findMedianSortedArrays2([]int{}, []int{1, 3, 4, 5}))
	fmt.Println(findMedianSortedArrays2([]int{1, 2, 3}, []int{}))
	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	//将偶数和奇数的情况合并，如果是奇数，会求两次同样的 k 。
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left)+getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	//让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}

	if k == 1 {
		if nums1[start1] < nums2[start2] {
			return nums1[start1]
		}
		return nums2[start2]
	}
	mini := len1
	if len1 > k/2 {
		mini = k / 2
	}
	minj := len2
	if len2 > k/2 {
		minj = k / 2
	}
	i := start1 + mini - 1
	j := start2 + minj - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-minj)
	}

	return getKth(nums1, i+1, end1, nums2, start2, end2, k-mini)
}

func findMedianSortedArrays2(nums1, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	lValue, rVlaue := 0, 0
	lPos, rPos := 0, 0
	for i := 0; i <= (l1+l2)/2; i++ {
		lValue = rVlaue
		if lPos < l1 && (rPos >= l2 || nums1[lPos] < nums2[rPos]) {
			rVlaue = nums1[lPos]
			lPos++
		} else {
			rVlaue = nums2[rPos]
			rPos++
		}
	}

	if (l1+l2)&1 == 0 {
		return float64(lValue+rVlaue) / 2.0
	}
	return float64(rVlaue)
}
