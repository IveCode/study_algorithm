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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空位置插入新节点
	if root == nil {
		return NewTreeNode(val)
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

func printlnBST(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printlnBST(root.Left)
	printlnBST(root.Right)
}

func isValidBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	if max != nil && root.Val > max.Val {
		return false
	}
	return isValidBST(root.Left, min, root) && isValidBST(root.Right, root, max)
}

func getMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		node := getMin(root.Right)
		root.Val = node.Val
		deleteNode(root.Right, node.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func main() {
	root := insertIntoBST(nil, 10)
	root = insertIntoBST(root, 9)
	root = insertIntoBST(root, 15)
	root = insertIntoBST(root, 13)
	root = insertIntoBST(root, 18)
	root = insertIntoBST(root, 16)
	printlnBST(root)
	fmt.Println()
	fmt.Println("isValidBST:", isValidBST(root, nil, nil))
	root = deleteNode(root, 15)
	printlnBST(root)
	fmt.Println()

	treeNode := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	data := serialize(treeNode)
	fmt.Println("serialize:", data)
	treeNode1 := deserialize(data)
	printlnBST(treeNode1)
	fmt.Println()
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
