package union_find

import (
	"fmt"
	"testing"
)

type UFQuickUnion struct {
	id    []int
	count int
}

func NewUFQuickUnion(n int) *UFQuickUnion {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &UFQuickUnion{
		id:    id,
		count: n,
	}
}

func (uf *UFQuickUnion)Count() int {
	return uf.count
}

func (uf *UFQuickUnion)Connected(p, q int) bool  {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UFQuickUnion)Find(p int) int  {
	for uf.id[p] != p {
		p = uf.id[p]
	}
	return p
}

func (uf *UFQuickUnion)Union(p , q int)   {
	//TODO:
	pId, qId := uf.Find(p), uf.Find(q)
	if pId == qId {
		return
	}
	uf.id[p] = q // 将一颗树(即一个组)变成另外一课树(即一个组)的子树
	uf.count--
}

func Test_QuickUnion(t *testing.T)  {
	ids := [][]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{6, 7},
	}
	uf := NewUFQuickUnion(len(ids))
	for i := 0; i < len(ids); i++ {
		p, q := ids[i][0], ids[i][1]
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
	}
	fmt.Println("Ids:", uf.id)
	fmt.Println("Result:", uf.Count())
}