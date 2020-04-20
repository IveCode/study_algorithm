package main

import "fmt"

func main() {
	//exec([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	exec([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}})
}
func exec(matrix [][]int) {
	//for i := 0; i < len(matrix); i++ {
	//	for j := 0; j < len(matrix[i]); j++ {
	//		if matrix[i][j] == 1 {
	//			matrix[i][j] = -1
	//		}
	//	}
	//}
	matrix = updateMatrix(matrix)
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println("===================")
}

func updateMatrix(matrix [][]int) [][]int {
	m := NewMatrix(matrix)
	m.Process()
	return m.matrix
}

var direct = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Matrix struct {
	matrix [][]int
	RowLen int
	ColLen int
}

func NewMatrix(matrix [][]int) *Matrix {
	return &Matrix{matrix: matrix, RowLen: len(matrix), ColLen: len(matrix[0])}
}

func (m *Matrix) Process() {
	var coords [][]int
	for i := 0; i < m.RowLen; i++ {
		for j := 0; j < m.ColLen; j++ {
			if m.matrix[i][j] == 0 {
				coords = append(coords, []int{i, j})
			} else {
				m.matrix[i][j] = -1
			}
		}
	}

	index := 1
	for len(coords) != 0 {
		tmpCoords := coords
		coords = [][]int{}
		set := map[string]struct{}{}
		for _, coord := range tmpCoords {
			for i := 0; i < len(direct); i++ {
				x := coord[0] + direct[i][0]
				y := coord[1] + direct[i][1]
				if x < 0 || y < 0 || x >= m.RowLen || y >= m.ColLen {
					continue
				}
				if m.matrix[x][y] == -1 {
					m.matrix[x][y] = index
					key := fmt.Sprintf("%d_%d", x, y)
					if _, ok := set[key]; !ok {
						coords = append(coords, []int{x, y})
						set[key] = struct{}{}
					}
				}
			}
		}
		index++
	}
}
