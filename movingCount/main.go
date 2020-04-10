package main

import "fmt"

func main() {
	//fmt.Println(movingCount(2, 3, 1))
	//fmt.Println(movingCount(3, 1, 0))
	//fmt.Println(movingCount(15, 15, 1))
	fmt.Println(movingCount(16, 8, 4))
	fmt.Println(movingCount2(16, 8, 4))
	fmt.Println(movingCount(38, 15, 9))
	fmt.Println(movingCount2(38, 15, 9))
}

func movingCount(m int, n int, k int) int {
	var result int
	flags := make([][]int, m)
	for i := 0; i < m; i++ {
		flags[i] = make([]int, n)
	}

	flags[0][0] = 1
	result++

	for i := 0; i < m; i++ {
		hi := i % 10
		li := i / 10
		for j := 0; j < n; j++ {
			hj := j % 10
			lj := j / 10
			if (i > 0 && flags[i-1][j] == 1) || (j > 0 && flags[i][j-1] == 1) {
				if hi+li+hj+lj <= k {
					flags[i][j] = 1
					result++
				}
			}

		}
	}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		fmt.Printf("%d ", flags[i][j])
	//	}
	//	fmt.Println()
	//}
	return result
}

func movingCount2(m int, n int, k int) int {
	result := 0
	queue := make([][]int, 0)
	direct := [][]int{{1, 0}, {0, 1}}
	queue = append(queue, []int{0, 0})
	set := map[string]struct{}{}
	for len(queue) != 0 {
		tempQueue := queue
		queue = make([][]int, 0)
		for i := 0; i < len(tempQueue); i++ {
			coords := tempQueue[i]
			for j := 0; j < len(direct); j++ {
				x := coords[0] + direct[j][0]
				y := coords[1] + direct[j][1]
				if x >= m || y >= n {
					continue
				}
				hx := x / 10
				lx := x % 10
				hy := y / 10
				ly := y % 10
				if hx+lx+hy+ly > k {
					continue
				}
				if _, ok := set[fmt.Sprintf("%d,%d", x, y)]; ok {
					continue
				}
				set[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
				queue = append(queue, []int{x, y})
				result++
			}
		}
	}
	return result + 1
}
