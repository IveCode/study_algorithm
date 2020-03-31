package main

//118. 杨辉三角 https://leetcode-cn.com/problems/pascals-triangle/

import "fmt"

func main() {
	//fmt.Println(generate(5))
	//fmt.Println(generate(30))

	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(27))
}

func generate(numRows int) [][]int {
	var results [][]int
	for i := 0; i < numRows; i++ {
		var result []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result = append(result, 1)
				continue
			}

			if len(results) >= 2 && j >= 1 {
				result = append(result, results[i-1][j-1]+results[i-1][j])
			}
		}
		results = append(results, result)
	}
	return results
}

func generate1(numRows int) [][]int {
	var results [][]int
	for i := 0; i < numRows; i++ {
		var result []int
		for j := 0; j <= i; j++ {
			result = append(result, PascalTriangle(i, j))
		}
		results = append(results, result)
	}
	return results
}

func PascalTriangle(i, j int) int {
	if i == 0 || j == 0 || i == j {
		return 1
	}

	return PascalTriangle(i-1, j-1) + PascalTriangle(i-1, j)
}

func getRow1(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for j := 0; j <= rowIndex; j++ {
		result[j] = PascalTriangle(rowIndex, j)
	}
	return result
}

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = 1
		for j := i; j > 1; j-- {
			result[j-1] = result[j-2] + result[j-1]
		}
	}
	return result
}
