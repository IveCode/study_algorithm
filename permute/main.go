package main

import "fmt"

//46. 全排列  https://leetcode-cn.com/problems/permutations/

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
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
