package union_find

import (
	"fmt"
	"testing"
)

type UFQuickFind struct {
	id    []int
	count int
}

func NewUFQuickFind(n int) *UFQuickFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &UFQuickFind{
		id:    id,
		count: n,
	}
}

func (uf *UFQuickFind) Union(p, q int) {
	pId, qId := uf.Find(p), uf.Find(q)
	if pId == qId {
		return
	}

	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pId {
			uf.id[i] = qId
		}
	}
	uf.count--
}

func (uf *UFQuickFind) Find(p int) int {
	return uf.id[p]
}

func (uf *UFQuickFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UFQuickFind) Count() int {
	return uf.count
}

func Test_QuickFind(t *testing.T) {
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
	uf := NewUFQuickFind(len(ids))
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
