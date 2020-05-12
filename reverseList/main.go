package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	printList(l)
	r := reverseList(l)
	printList(r)
}

func printList(head *ListNode) {
	p := head
	for ; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println("nil")
}

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
