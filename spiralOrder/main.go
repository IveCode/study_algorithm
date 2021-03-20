package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             //[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) //[1,2,3,4,8,12,11,10,9,5,6,7]
}

var direction = []int{1, 1, -1, -1}

func spiralOrder(matrix [][]int) []int {
	m1, m2 := 0, len(matrix)
	n1, n2 := 0, len(matrix[0])
	length := m2 * n2
	result := make([]int, 0, length)

	i, j := m1, n1
	isRow := false
	idx := 0
	for k := 0; k < length; k++ {
		result = append(result, matrix[i][j])
		if j == n2-1 && !isRow && direction[idx] == 1 {
			m1++
			isRow = true
			idx++
			idx %= len(direction)
		}
		if i == m2-1 && isRow && direction[idx] == 1 {
			n2--
			isRow = false
			idx++
			idx %= len(direction)
		}
		if j == n1 && !isRow && direction[idx] == -1 {
			m2--
			isRow = true
			idx++
			idx %= len(direction)
		}
		if i == m1 && isRow && direction[idx] == -1 {
			n1++
			isRow = false
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
