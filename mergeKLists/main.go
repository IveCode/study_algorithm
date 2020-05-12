package main

import "fmt"

/*
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func main() {
	list := mergeKLists([]*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
	})
	printList(list)
}

func printList(list *ListNode) {
	p := list
	for {
		if p != nil && p.Next != nil {
			fmt.Printf("%d->", p.Val)
		}
		if p != nil && p.Next == nil {
			fmt.Printf("%d", p.Val)
			break
		}
		p = p.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := &ListNode{}
	p := l
	var min *ListNode
	var index int
	for {
		min = nil
		index = 0
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if min == nil || min.Val > lists[i].Val {
				min = lists[i]
				index = i
			}
		}

		if min == nil || index == -1 {
			break
		}
		p.Next = &ListNode{Val: min.Val}
		p = p.Next

		lists[index] = lists[index].Next
	}

	p = l.Next
	l.Next = nil

	return p
}
