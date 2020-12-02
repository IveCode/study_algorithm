package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

type Node struct {
	Name string
}

func (n Node) String() string {
	return n.Name
}

type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

type Graph struct {
	Edges []*Edge
	Nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{
		Edges: []*Edge{},
		Nodes: []*Node{},
	}
}

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
		s.WriteString(fmt.Sprintf("\t%s -> %s = %d\n", edge.Parent, edge.Child, edge.Cost))
	}

	s.WriteString("Node:\n\t")
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s.WriteString(fmt.Sprintf("%s\n", node))
		} else {
			s.WriteString(fmt.Sprintf("%s, ", node))
		}
	}
	return s.String()
}

func (g *Graph) Dijkstra(startNode *Node) {
	costTable := g.NewCostTable(startNode)

	visited := []*Node{}

	for len(visited) != len(g.Nodes) {
		node := getClosesNonVisitedNode(costTable, visited)

		visited = append(visited, node)
		nodeEdges := g.GetNodeEdges(node)

		for _, edge := range nodeEdges {
			distanceToNeighbor := costTable[node] + edge.Cost
			if distanceToNeighbor < costTable[edge.Child] {
				costTable[edge.Child] = distanceToNeighbor
			}
		}
	}

	fmt.Println("Result:")
	for node, cost := range costTable {
		fmt.Println("	", startNode, "->", node, "=", cost)
	}
}

func (g *Graph) NewCostTable(startNode *Node) map[*Node]int {
	costTable := map[*Node]int{
		startNode: 0,
	}
	for _, node := range g.Nodes {
		if node.Name == startNode.Name {
			continue
		}
		costTable[node] = math.MaxInt32
	}
	return costTable
}

func (g *Graph) GetNodeEdges(node *Node) []*Edge {
	var edges []*Edge
	for i, edge := range g.Edges {
		if edge.Parent.Name == node.Name {
			edges = append(edges, g.Edges[i])
		}
	}
	return edges
}

func getClosesNonVisitedNode(costTable map[*Node]int, visited []*Node) *Node {
	type CostTableSort struct {
		Node *Node
		Cost int
	}
	var cts []CostTableSort
	for node, cost := range costTable {
		isVisited := false
		for _, visitedNode := range visited {
			if node.Name == visitedNode.Name {
				isVisited = true
				break
			}
		}
		if !isVisited {
			cts = append(cts, CostTableSort{
				Node: node,
				Cost: cost,
			})
		}
	}

	sort.Slice(cts, func(i, j int) bool {
		return cts[i].Cost < cts[j].Cost
	})
	return cts[0].Node
}

func main() {
	A := &Node{Name: "A"}
	B := &Node{Name: "B"}
	C := &Node{Name: "C"}
	D := &Node{Name: "D"}

	graph := NewGraph()
	graph.AddEdge(A, B, 6)
	graph.AddEdge(A, C, 3)
	graph.AddEdge(C, B, 2)
	graph.AddEdge(B, D, 1)
	graph.AddEdge(C, D, 5)

	fmt.Printf("%s", graph.String())
	graph.Dijkstra(A)
}
