package main

import (
	"fmt"
)

func main() {
	PrintResult([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}

func PrintResult(board [][]byte) {
	solve(board)
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
}

type Coordinate struct {
	row, col int
}

func solve(board [][]byte) {
	var queue []Coordinate
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, Coordinate{i, 0})
		}
		if board[i][col-1] == 'O' {
			queue = append(queue, Coordinate{i, col - 1})
		}
	}

	for i := 1; i < col-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, Coordinate{0, i})
		}
		if board[row-1][i] == 'O' {
			queue = append(queue, Coordinate{row - 1, i})
		}
	}

	filter := map[string]struct{}{}
	for i := 0; i < len(queue); i++ {
		coordinate := queue[i]
		key := fmt.Sprintf("%d_%d", coordinate.row, coordinate.col)
		if _, ok := filter[key]; ok {
			continue
		}
		filter[key] = struct{}{}
		if coordinate.row-1 > 0 && board[coordinate.row-1][coordinate.col] == 'O' {
			queue = append(queue, Coordinate{coordinate.row - 1, coordinate.col})
		}
		if coordinate.row+1 < row && board[coordinate.row+1][coordinate.col] == 'O' {
			queue = append(queue, Coordinate{coordinate.row + 1, coordinate.col})
		}
		if coordinate.col-1 > 0 && board[coordinate.row][coordinate.col-1] == 'O' {
			queue = append(queue, Coordinate{coordinate.row, coordinate.col - 1})
		}
		if coordinate.col+1 < col && board[coordinate.row][coordinate.col+1] == 'O' {
			queue = append(queue, Coordinate{coordinate.row, coordinate.col + 1})
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				key := fmt.Sprintf("%d_%d", i, j)
				if _, ok := filter[key]; !ok {
					board[i][j] = 'X'
				}
			}
		}
	}
}
