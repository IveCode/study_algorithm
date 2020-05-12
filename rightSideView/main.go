package main

import "fmt"

//199. 二叉树的右视图 https://leetcode-cn.com/problems/binary-tree-right-side-view/

func main() {
	rsv(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	})

	rsv(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	})
}

func rsv(root *TreeNode) {
	fmt.Println(rightSideView(root))
	fmt.Println(rightSideView2(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rsv := NewRightSideView()
	rsv.dfs(root, 0)
	return rsv.Result
}

type RightSideView struct {
	Result []int
}

func NewRightSideView() *RightSideView {
	return &RightSideView{Result: []int{}}
}

func (r *RightSideView) dfs(root *TreeNode, deep int) {
	if root == nil {
		return
	}

	if deep == len(r.Result) {
		r.Result = append(r.Result, root.Val)
	}

	//fmt.Println(deep, r.Result)
	r.dfs(root.Right, deep+1)
	r.dfs(root.Left, deep+1)
}

//方法二 bfs
func rightSideView2(root *TreeNode) []int {
	var result []int
	var deep int

	queue := []*TreeNode{root}
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			if deep == len(result) {
				result = append(result, queue[0].Val)
			}

			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			length--
			queue = queue[1:]
		}
		deep++
	}

	return result
}
