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
	data := "4,2,1,0,null,null,null,3,null,null,5,null,6,null,null,"
	root := deserialize(data)
	fmt.Println(serialize(root))
	fmt.Println(data)
	root2 := convertBiNode(root)
	fmt.Println(serialize(root2))
}

func convertBiNode(root *TreeNode) *TreeNode {
	cur := &TreeNode{}
	build(root, cur)
	return cur.Right
}

func build(root, cur *TreeNode) *TreeNode {
	if root == nil {
		return cur
	}
	cur = build(root.Left, cur)
	root.Left = nil
	cur.Right = root
	cur = root
	cur = build(root.Right, cur)
	return cur
}
