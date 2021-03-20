package main

import "fmt"

//https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	var reverse bool
	for len(queue) != 0 {
		length := len(queue)
		values := make([]int, length)
		reverse = !reverse
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if reverse {
				values[i] = queue[i].Val
			} else {
				values[length-i-1] = queue[i].Val
			}
		}
		queue = queue[length:]
		result = append(result, values)
	}

	return result
}
