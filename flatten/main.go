package main

import "fmt"

//https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
//114. 二叉树展开为链表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: &TreeNode{Val: 6},
		},
	}

	flatten(root)
	printlnTreeNode(root)
}

func printlnTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	//printlnTreeNode(root.Left)
	printlnTreeNode(root.Right)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表,将左子树作为右子树
	right := root.Right
	root.Right = root.Left
	root.Left = nil

	//2、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
