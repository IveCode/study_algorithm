package lc_test

import (
	"testing"
)

func TestMerge(t *testing.T) {
	printMerge(t, []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	printMerge(t, []int{1, 2, 3, 0, 0, 0}, 3, []int{1, 1, 1}, 3)
	printMerge(t, []int{5, 6, 7, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	printMerge(t, []int{5, 0}, 1, []int{1}, 1)
}

func printMerge(t *testing.T, nums1 []int, m int, nums2 []int, n int) {
	t.Log("nums1:", nums1, "m:", m)
	t.Log("nums2:", nums2, "n:", n)
	merge2(nums1, m, nums2, n)
	t.Log("results:", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; i < m+n && j < n; {
		if i < m+j && nums1[i] <= nums2[j] {
			i++
			continue
		} else if i > m+j {
			nums1[i] = nums2[j]
			i++
			j++
		} else {
			for k := m + j; k > i; k-- {
				nums1[k] = nums1[k-1]
			}
			nums1[i] = nums2[j]
			i++
			j++
		}
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	k, i, j := m+n-1, m-1, n-1

	for k >= 0 {
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || (j < 0 && i >= 0) {
			nums1[k] = nums1[i]
			k--
			i--
		}

		if (i >= 0 && j >= 0 && nums1[i] <= nums2[j]) || (i < 0 && j >= 0) {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

}
