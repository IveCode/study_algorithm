package main

import (
	"fmt"
)

//面试题64. 求1+2+…+n https://leetcode-cn.com/problems/qiu-12n-lcof/

func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(9))
}

func sumNums(n int) int {
	s := &Sum{}
	s.SumNums(n)
	return s.result
}

type Sum struct {
	result int
}

func (this *Sum) SumNums(n int) bool {
	this.result += n
	return n != 0 && this.SumNums(n-1)
}
