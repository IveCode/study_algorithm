package main

import (
	"fmt"
)

//983. 最低票价 https://leetcode-cn.com/problems/minimum-cost-for-tickets/

func main() {
	printMinCostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})                                                       //11
	printMinCostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})                                   //17
	printMinCostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15})                                                       //6
	printMinCostTickets([]int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28}, []int{3, 13, 45}) //44
}

func printMinCostTickets(days, costs []int) {
	fmt.Println("days:", days, "costs:", costs, "result", mincostTickets(days, costs))
}

func mincostTickets(days []int, costs []int) int {
	money = make([]int, len(days))
	return costTicket(days, costs, 0)
}

var money []int

var values = []int{1, 7, 30}

func costTicket(days, costs []int, dayIndex int) int {
	if dayIndex >= len(days) {
		return 0
	}

	if money[dayIndex] != 0 {
		return money[dayIndex]
	}

	j := dayIndex
	for i := 0; i < len(values); i++ {
		for j < len(days) && days[j] < days[dayIndex]+values[i] {
			j++
		}

		result := costTicket(days, costs, j) + costs[i]

		if money[dayIndex] == 0 || money[dayIndex] > result {
			money[dayIndex] = result
		}
	}

	return money[dayIndex]
}
