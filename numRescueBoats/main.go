package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var result int
	for i, j := 0, len(people)-1; i <= j; {
		result++
		if people[i]+people[j] <= limit {
			i++
		}
		j--
	}
	return result
}
