package main

import "fmt"

/*
63. 不同路径 II
	https://leetcode-cn.com/problems/unique-paths-ii/
*/
func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	x, y := len(obstacleGrid), len(obstacleGrid[0])
	results := make([]int, y)
	if obstacleGrid[0][0] == 0 {
		results[0] = 1
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if obstacleGrid[i][j] == 1 {
				results[j] = 0
				continue
			}

			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				results[j] += results[j-1]
			}
		}
	}
	fmt.Println(results)
	return results[len(results)-1]
}
