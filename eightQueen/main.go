package main

import "fmt"

//八皇后问题 => n*n
func main() {
	N := 4
	queen := NewQueen(N)
	queen.DFS(0)
	fmt.Println("===============")
}

type Queen struct {
	Result [][]int
	N      int
}

func NewQueen(n int) *Queen {
	queen := &Queen{}
	queen.Result = make([][]int, n)
	for i := 0; i < n; i++ {
		queen.Result[i] = make([]int, n)
	}
	queen.N = n
	return queen
}

func (q *Queen) DFS(i int) {

	if i == q.N {
		q.Print()
		return
	}

	for j := 0; j < q.N; j++ {
		if q.Check(i, j) {
			q.Result[i][j] = 1
			q.DFS(i + 1)
			q.Result[i][j] = 0
		}
	}
}

func (q *Queen) Print() {
	for i := 0; i < len(q.Result); i++ {
		for j := 0; j < len(q.Result); j++ {
			fmt.Printf("%d ", q.Result[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==============")
}

func (q *Queen) Check(x, y int) bool {
	for i := 0; i < q.N; i++ {
		if q.Result[i][y] == 1 {
			return false
		}
	}

	for i := 0; x-i >= 0 && y-i >= 0; i++ {
		if q.Result[x-i][y-i] == 1 {
			return false
		}
	}

	for i := 0; x-i > 0 && y+i < q.N; i++ {
		if q.Result[x-i][y+i] == 1 {
			return false
		}
	}

	return true
}
