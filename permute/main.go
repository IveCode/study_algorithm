package main

import "fmt"

//46. 全排列  https://leetcode-cn.com/problems/permutations/

func main() {
	//fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute2([]int{5, 4, 6, 2}))
}

func permute(nums []int) [][]int {
	p := NewPermute()
	p.permute(nums, 0, len(nums)-1)
	return p.Result
}

type Permute struct {
	Result [][]int
}

func NewPermute() *Permute {
	return &Permute{Result: [][]int{}}
}

func (p *Permute) permute(nums []int, k, m int) {
	if k == m {
		temp := make([]int, len(nums))
		copy(temp, nums)
		p.Result = append(p.Result, temp)
		return
	}

	for i := k; i <= m; i++ {
		nums[i], nums[k] = nums[k], nums[i]
		p.permute(nums, k+1, m)
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func permute2(nums []int) [][]int {
	var results [][]int
	var backtrack func([]int, []int)

	backtrack = func(nums, result []int) {
		if len(result) == len(nums) {
			fmt.Println(result)
			temp := make([]int, len(result))
			copy(temp, result)
			results = append(results, temp)
			fmt.Println(results)
			return
		}

		for i := 0; i < len(nums); i++ {
			isContain := false
			for _, v := range result {
				if v == nums[i] {
					isContain = true
					break
				}
			}
			if isContain {
				continue
			}
			result = append(result, nums[i])
			backtrack(nums, result)
			result = result[:len(result)-1]
		}
	}

	var result []int
	backtrack(nums, result)
	return results
}
