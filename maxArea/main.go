package main

import "fmt"

//盛最多水的容器 leecode:https://leetcode-cn.com/problems/container-with-most-water/
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	for l < r {
		min := height[l]
		if min > height[r] {
			min = height[r]
		}

		tempArea := (r - l) * min
		if area < tempArea {
			area = tempArea
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func maxArea2(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if min > height[j] {
				min = height[j]
			}

			tmpArea := (j - i) * min
			if area < tmpArea {
				area = tmpArea
			}
		}
	}
	return area
}
