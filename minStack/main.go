package main

import "fmt"

//155. 最小栈 https://leetcode-cn.com/problems/min-stack
func main() {
	minStack := Constructor()
	minStack.Pop()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // --> 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top())    //--> 返回 0.
	fmt.Println(minStack.GetMin()) //--> 返回 -2.

}

type MinStack struct {
	value    []int
	length   int
	minValue []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		value:    []int{},
		minValue: []int{},
		length:   0,
	}
}

func (this *MinStack) Push(x int) {
	this.value = append(this.value, x)
	if this.length == 0 || x < this.minValue[this.length-1] {
		this.minValue = append(this.minValue, x)
	} else {
		this.minValue = append(this.minValue, this.minValue[this.length-1])
	}
	this.length++
}

func (this *MinStack) Pop() {
	if this.length > 0 {
		this.length--
		this.value = this.value[:this.length]
		this.minValue = this.minValue[:this.length]
	}
}

func (this *MinStack) Top() int {
	if this.length > 0 {
		return this.value[this.length-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.length > 0 {
		return this.minValue[this.length-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
