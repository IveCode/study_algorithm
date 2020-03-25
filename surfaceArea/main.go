package main

import "fmt"

//892. 三维形体的表面积 https://leetcode-cn.com/problems/surface-area-of-3d-shapes/
func main() {
	fmt.Println(surfaceArea([][]int{[]int{2}}))
	fmt.Println(surfaceArea([][]int{[]int{1, 2}, []int{3, 4}}))
	fmt.Println(surfaceArea([][]int{[]int{1, 0}, []int{0, 2}}))
	fmt.Println(surfaceArea([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}))
	fmt.Println(surfaceArea([][]int{[]int{2, 2, 2}, []int{2, 1, 2}, []int{2, 2, 2}}))
}

func surfaceArea(grid [][]int) int {
	if grid == nil {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	area := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			area += grid[i][j] * 6
			if grid[i][j] > 1 { //高度
				area = area - (grid[i][j]-1)*2
			}
			if j > 0 {
				if grid[i][j] > grid[i][j-1] {
					area = area - grid[i][j-1]*2
				} else {
					area = area - grid[i][j]*2
				}
			}
			if i > 0 {
				if grid[i][j] > grid[i-1][j] {
					area = area - grid[i-1][j]*2
				} else {
					area = area - grid[i][j]*2
				}
			}
		}
	}
	return area
}
