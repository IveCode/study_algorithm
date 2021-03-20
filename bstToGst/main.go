package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "null,"
	}
	str := fmt.Sprintf("%d,", root.Val)
	str += serialize(root.Left)
	str += serialize(root.Right)
	return str
}

func deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	values := strings.Split(data, ",")
	values = values[:len(values)-1]
	if len(values) == 0 {
		return nil
	}
	var build func() *TreeNode

	build = func() *TreeNode {
		value := values[0]
		values = values[1:]
		if value == "null" {
			return nil
		}
		v, _ := strconv.ParseInt(value, 10, 64)
		root := NewTreeNode(int(v))
		root.Left = build()
		root.Right = build()
		return root
	}

	return build()
}

func main() {
	data := "4,1,0,null,null,2,null,3,null,null,6,5,null,null,7,null,8,null,null,"
	root := deserialize(data)
	root1 := deserialize(data)
	//fmt.Println(data == serialize(root))
	fmt.Println(serialize(bstToGst1(root)))
	fmt.Println(serialize(bstToGst(root1)))
}
func bstToGst(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	sum := 0
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}

func bstToGst1(root *TreeNode) *TreeNode {
	total := preorderBST(root)
	preorderBSTToGst(root, total)
	return root
}

func preorderBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	sum += preorderBST(root.Left)
	sum += root.Val
	sum += preorderBST(root.Right)
	return sum
}

func preorderBSTToGst(root *TreeNode, value int) int {
	if root == nil {
		return value
	}
	value = preorderBSTToGst(root.Left, value)
	val := root.Val
	root.Val = value
	value -= val
	value = preorderBSTToGst(root.Right, value)
	return value
}
