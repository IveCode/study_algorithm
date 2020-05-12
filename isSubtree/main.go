package main

import "fmt"

func main() {
	fmt.Println(isSubtree(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}, &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}

	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	return s != nil && t != nil && s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
