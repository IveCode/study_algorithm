package main

import "fmt"

//knapsack problem

func main() {
	fmt.Println(knapsack([][]int{{2, 3}, {3, 4}, {4, 5}, {5, 8}, {9, 10}}, 20))
}

//goods {weight, value}
func knapsack(goods [][]int, total int) int {
	goods = append([][]int{{0, 0}}, goods...)
	result := make([][]int, len(goods))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, total+1)
	}

	for i := 1; i < len(goods); i++ {
		for j := 1; j <= total; j++ {
			weight, value := goods[i][0], goods[i][1]
			if weight > j {
				result[i][j] = result[i-1][j]
			} else {
				value1 := result[i-1][j-weight] + value
				value2 := result[i-1][j]
				if value1 > value2 {
					result[i][j] = value1
				} else {
					result[i][j] = value2
				}
			}
		}
	}

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

	return result[len(result)-1][total]
}
