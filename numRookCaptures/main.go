package main

import "fmt"

//999. 车的可用捕获量 https://leetcode-cn.com/problems/available-captures-for-rook/

func main() {
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'R', '.', '.', '.', 'p'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}))
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'B', 'R', 'B', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}))
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'p', 'p', '.', 'R', '.', 'p', 'B', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}))
}

func numRookCaptures(board [][]byte) int {
	p := 0
	x, y := findR(board)
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'B' {
			break
		}
		if board[i][y] == '.' {
			continue
		}
		if board[i][y] == 'p' {
			p++
			break
		}
	}

	for i := x + 1; i < len(board); i++ {
		if board[i][y] == 'B' {
			break
		}
		if board[i][y] == '.' {
			continue
		}
		if board[i][y] == 'p' {
			p++
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		if board[x][i] == 'B' {
			break
		}
		if board[x][i] == '.' {
			continue
		}
		if board[x][i] == 'p' {
			p++
			break
		}
	}

	for i := y + 1; i < len(board[x]); i++ {
		if board[x][i] == 'B' {
			break
		}
		if board[x][i] == '.' {
			continue
		}
		if board[x][i] == 'p' {
			p++
			break
		}
	}
	return p
}

func findR(board [][]byte) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}
	return 0, 0
}
