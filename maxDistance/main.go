package main

//1162. 地图分析 https://leetcode-cn.com/problems/as-far-from-land-as-possible/

import "fmt"

func main() {
	fmt.Println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
	fmt.Println(maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

type Land struct {
	x, y int
}

func maxDistance(grid [][]int) int {
	N := len(grid)
	var lands []Land
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				lands = append(lands, Land{
					x: i,
					y: j,
				})
			}
		}
	}

	result := -1
	if len(lands) == 0 || len(lands) == N*N {
		return result
	}

	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(lands) != 0 {
		result++
		allLands := lands
		lands = []Land{}
		for i := 0; i < len(allLands); i++ {
			set := map[Land]struct{}{}
			for j := 0; j < len(direction); j++ {
				x := allLands[i].x + direction[j][0]
				y := allLands[i].y + direction[j][1]
				if x < 0 || y < 0 || x >= N || y >= N {
					continue
				}
				if grid[x][y] == 0 {
					grid[x][y] = 1
					land := Land{
						x: x,
						y: y,
					}
					if _, ok := set[land]; !ok {
						set[land] = struct{}{}
					}
				}
			}

			for k, _ := range set {
				lands = append(lands, k)
			}
		}
	}

	return result
}
