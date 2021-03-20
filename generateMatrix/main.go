package main

import "fmt"

func main() {
	//fmt.Println(generateMatrix2(5))
	fmt.Println(generateMatrix2(3))
	//fmt.Println(generateMatrix2(2))
	//fmt.Println(generateMatrix2(1))
}

func generateMatrix2(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx := 0
	x, y := 0, 0

	for i := 0; i < n*n; i++ {
		result[x][y] = i + 1
		dir := direction[idx]
		if nextX, nextY := x+dir[0], y+dir[1]; nextX == -1 || nextX == n || nextY == -1 || nextY == n || result[nextX][nextY] > 0 {
			idx++
			idx %= len(direction)
			dir = direction[idx]
		}
		x += dir[0]
		y += dir[1]
	}

	return result
}

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	direction := []int{1, 1, -1, -1}
	idx := 0
	m1, m2 := 0, n-1 //row
	n1, n2 := 0, n-1 //col
	isRow := false
	i, j := 0, 0
	var isChange bool
	for k := 1; k <= n*n; k++ {
		result[i][j] = k
		isChange = false
		if j == n2 && direction[idx] == 1 && !isRow {
			m1++
			isChange = true
		}
		if i == m2 && direction[idx] == 1 && isRow {
			n2--
			isChange = true
		}
		if j == n1 && direction[idx] == -1 && !isRow {
			m2--
			isChange = true
		}

		if i == m1 && direction[idx] == -1 && isRow {
			n1++
			isChange = true
		}

		if isChange {
			isRow = !isRow
			idx++
			idx %= len(direction)
		}

		if isRow {
			i += direction[idx]
		} else {
			j += direction[idx]
		}
	}
	return result
}
