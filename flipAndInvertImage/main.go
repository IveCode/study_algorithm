package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
func flipAndInvertImage(A [][]int) [][]int {
	m, n := len(A), len(A[0])-1
	for i := 0; i < m; i++ {
		for j, k := 0, n; j <= k; {
			if j != k {
				A[i][j], A[i][k] = reversal(A[i][k]), reversal(A[i][j])
			} else {
				A[i][j] = reversal(A[i][k])
			}
			j++
			k--
		}
	}
	_ = n
	return A
}
func reversal(v int) int {
	if v == 1 {
		return 0
	}
	return 1
}
