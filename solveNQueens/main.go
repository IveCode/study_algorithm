package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var results [][]string

	var backtrack func([][]byte, int)

	backtrack = func(result [][]byte, row int) {
		if row == n {
			var str []string
			for i := 0; i < n; i++ {
				str = append(str, string(result[i]))
			}
			results = append(results, str)
			return
		}

		for j := 0; j < n; j++ {
			if isValid(result, row, j, n) {
				result[row][j] = 'Q'
				backtrack(result, row+1)
				result[row][j] = '.'
			}
		}
	}

	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = '.'
		}
	}

	backtrack(result, 0)

	return results
}

func CountNQueens(result [][]byte, n int) bool {
	num := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if result[i][j] == 'Q' {
				num++
			}
		}
	}
	return num == n
}

func isValid(result [][]byte, row, col, n int) bool {
	for i := 0; i < n; i++ {
		if result[i][col] == 'Q' {
			return false
		}
	}

	for j := 0; j < n; j++ {
		if result[row][j] == 'Q' {
			return false
		}
	}

	for i := 0; row-i >= 0 && col-i >= 0; i++ {
		if result[row-i][col-i] == 'Q' {
			return false
		}
	}
	for i := 0; row-i >= 0 && col+i < n; i++ {
		if result[row-i][col+i] == 'Q' {
			return false
		}
	}
	return true
}
