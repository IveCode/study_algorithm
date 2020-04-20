package main

import (
	"fmt"
	"sort"
)

//56. 合并区间 https://leetcode-cn.com/problems/merge-intervals/

func main() {
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(ArrayIntSort(intervals))
	var results [][]int

	results = append(results, intervals[0])
	for i := 1; i < len(intervals); i++ {
		idx := len(results) - 1
		if results[idx][1] >= intervals[i][0] {
			if results[idx][1] < intervals[i][1] {
				results[idx][1] = intervals[i][1]
			}
			continue
		}
		results = append(results, intervals[i])
	}
	return results
}

type ArrayIntSort [][]int

func (p ArrayIntSort) Len() int           { return len(p) }
func (p ArrayIntSort) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p ArrayIntSort) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
