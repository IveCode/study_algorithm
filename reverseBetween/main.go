package main

import "fmt"

func main() {
	Exec([]int{1, 2, 3, 4, 5}, 2, 4)
	Exec([]int{5}, 1, 1)
	Exec([]int{3, 5}, 1, 2)
}

func Exec(nums []int, left, right int) {
	root := NewListNode(nums)
	fmt.Printf("init:")
	printListNode(root)
	fmt.Printf("result:")
	printListNode(reverseBetween(root, left, right))
}

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: NewListNode(nums[1:]),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(root *ListNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	if root.Next != nil {
		fmt.Printf("->")
		printListNode(root.Next)
	} else {
		fmt.Printf("\n")
	}

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	vmHead := &ListNode{
		Val:  0,
		Next: head,
	}
	leftPreNode := vmHead
	for i := 0; i < left-1; i++ {
		leftPreNode = leftPreNode.Next
	}

	curNode := leftPreNode.Next
	for i := 0; i < right-left; i++ {
		nextNode := curNode.Next
		curNode.Next = nextNode.Next
		nextNode.Next = leftPreNode.Next
		leftPreNode.Next = nextNode
	}
	return vmHead.Next
}
