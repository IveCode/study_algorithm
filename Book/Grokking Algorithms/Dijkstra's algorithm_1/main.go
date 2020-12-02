package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

//Node Name
type Node struct {
	Name string
}

func (n Node) String() string {
	return n.Name
}

//Edge is `A-(x)->B`, Parent is A, Child is B, Cost X
type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

//Graph 是一个带权有向图
type Graph struct {
	Edges []*Edge
	Nodes []*Node
}

//NewGraph 生成一个Graph对象
func NewGraph() *Graph {
	return &Graph{
		Edges: []*Edge{},
		Nodes: []*Node{},
	}
}

//AddEdge 添加边
func (g *Graph) AddEdge(parent, child *Node, cost int) {
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

//AddNode 添加节点
func (g *Graph) AddNode(node *Node) {
	for _, n := range g.Nodes {
		if n.Name == node.Name {
			return
		}
	}
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) String() string {
	s := bytes.NewBufferString("Edges:\n")

	for _, edge := range g.Edges {
		s.WriteString(fmt.Sprintf("\t%s -> %s = %d\n", edge.Parent.Name, edge.Child.Name, edge.Cost))
	}

	s.WriteString("Nodes:\n\t")
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s.WriteString(fmt.Sprintf("%s\n", node.Name))
		} else {
			s.WriteString(fmt.Sprintf("%s, ", node.Name))
		}
	}

	return s.String()
}

//Dijkstra 算法
func (g *Graph) Dijkstra(startNode *Node) {
	costTable := g.NewCostTable(startNode)
	var arrived []*Node

	for len(arrived) != len(g.Nodes) {
		node := FindNextNode(costTable, arrived)

		arrived = append(arrived, node)
		edges := g.FindEdgesNode(node)

		for _, edge := range edges {
			value := edge.Cost + costTable[edge.Parent]
			if value < costTable[edge.Child] {
				costTable[edge.Child] = value
			}
		}
	}

	fmt.Println("Result:")
	for node, cost := range costTable {
		fmt.Printf("\t%s -> %s = %d\n", startNode.Name, node.Name, cost)
	}
}

//NewCostTable 生成消耗记录表
func (g *Graph) NewCostTable(startNode *Node) map[*Node]int {
	costTable := map[*Node]int{
		startNode: 0,
	}

	for i, node := range g.Nodes {
		if node.Name == startNode.Name {
			continue
		}
		costTable[g.Nodes[i]] = math.MaxInt32
	}
	return costTable
}

//FindEdgesNode 查找节点关联的边
func (g *Graph) FindEdgesNode(parent *Node) []*Edge {
	var edges []*Edge
	for i, edge := range g.Edges {
		if edge.Parent.Name == parent.Name {
			edges = append(edges, g.Edges[i])
		}
	}
	return edges
}

//SortNode 用于排序
type SortNode struct {
	Node *Node
	Cost int
}

//FindNextNode 查找最优的节点
func FindNextNode(costTable map[*Node]int, arrived []*Node) *Node {
	var sortNode []SortNode

	for node, cost := range costTable {
		isArrived := false
		for _, arrivedNode := range arrived {
			if arrivedNode.Name == node.Name {
				isArrived = true
				break
			}
		}
		if !isArrived {
			sortNode = append(sortNode, SortNode{
				Node: node,
				Cost: cost,
			})
		}
	}

	sort.Slice(sortNode, func(i, j int) bool {
		return sortNode[i].Cost < sortNode[j].Cost
	})

	return sortNode[0].Node
}

func main() {
	A := &Node{Name: "A"}
	B := &Node{Name: "B"}
	C := &Node{Name: "C"}
	D := &Node{Name: "D"}

	g := NewGraph()
	g.AddEdge(A, B, 6)
	g.AddEdge(A, C, 3)
	g.AddEdge(B, D, 1)
	g.AddEdge(C, B, 2)
	g.AddEdge(C, D, 5)

	fmt.Println(g)
	g.Dijkstra(A)
}
