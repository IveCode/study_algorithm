package main

import "fmt"

//Z 字形变换 leecode: https://leetcode-cn.com/problems/zigzag-conversion/
func main() {
	fmt.Println(convert("AB", 1))
}

func convert(s string, numRows int) string {
	if numRows < 1 {
		return ""
	}

	if numRows == 1 {
		return s
	}

	temp := make([]string, numRows)
	col := 0
	direction := false
	for i := 0; i < len(s); i++ {
		temp[col] += string(s[i])
		if col == 0 || col == numRows-1 {
			direction = !direction
		}
		if direction {
			col++
		} else {
			col--
		}
	}

	result := ""
	for i := 0; i < len(temp); i++ {
		result += temp[i]
	}

	return result
}
