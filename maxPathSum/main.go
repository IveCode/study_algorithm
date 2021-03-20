package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxPathSum(ArrayToTree("1,2,3", 0)))
	fmt.Println(maxPathSum(ArrayToTree("-10,9,20,null,null,15,7", 0)))
	fmt.Println(maxPathSum(ArrayToTree("-10", 0)))
	//printBST(ArrayToTree("-10,9,20,null,null,15,7", 0))
}
func printBST(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printBST(root.Left)
	printBST(root.Right)
}

func ArrayToTree(data string, index int) *TreeNode {
	values := strings.Split(data, ",")
	if len(values) == 0 || index >= len(values) {
		return nil
	}
	if values[index] == "null" {
		return nil
	}
	v, _ := strconv.Atoi(values[index])
	root := NewTreeNode(v)

	if index*2+1 < len(values) {
		root.Left = ArrayToTree(data, index*2+1)
	}

	if index*2+2 < len(values) {
		root.Right = ArrayToTree(data, index*2+2)
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(0, maxGain(node.Left))
		rightSum := max(0, maxGain(node.Right))
		maxSum = max(maxSum, leftSum+rightSum+node.Val)

		return max(leftSum, rightSum) + node.Val
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
