package main

import "fmt"

//200. 岛屿数量 https://leetcode-cn.com/problems/number-of-islands/

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	g := NewGrid(grid)
	g.numIslands()
	return g.total
}

type Grid struct {
	grid  [][]byte
	rows  int
	cols  int
	total int
}

func NewGrid(grid [][]byte) *Grid {
	return &Grid{grid: grid, rows: len(grid), cols: len(grid[0])}
}

func (g *Grid) numIslands() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.grid[i][j] == '1' {
				g.total++
				g.dfs(i, j)
			}
		}
	}
}

var direct = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (g *Grid) dfs(x, y int) {
	g.grid[x][y] = '0'

	for i := 0; i < len(direct); i++ {
		tx := x + direct[i][0]
		ty := y + direct[i][1]
		if tx < 0 || ty < 0 || tx >= g.rows || ty >= g.cols {
			continue
		}
		if g.grid[tx][ty] == '1' {
			g.dfs(tx, ty)
		}
	}
}
