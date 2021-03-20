package Sort

import (
	"fmt"
	"testing"
)

type MergeSort struct {
}

func (s *MergeSort) Sort(a []int) {
	//s.merge(a, 0, len(a)/2, len(a)-1)
	s.TopDownMergeSort(a, 0, len(a)-1)
}

func (s *MergeSort) TopDownMergeSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (hi + lo) / 2
	s.TopDownMergeSort(a, lo, mid)
	s.TopDownMergeSort(a, mid+1, hi)
	s.merge(a, lo, mid, hi)
}

func (s *MergeSort) merge(a []int, lo, mid, hi int) {
	aux := make([]int, len(a))
	for i := lo; i < len(a); i++ {
		aux[i] = a[i]
	}

	min := func(a, b int) bool { return a < b }

	for k, i, j := lo, lo, mid+1; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if min(aux[i], aux[j]) {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}

	fmt.Println(lo, mid, hi, a)
}

func (s *MergeSort) ButtonUpMergeSort(a []int) {
	length := len(a)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for sz := 1; sz < length; sz = sz + sz {
		for lo := 0; lo < length-sz; lo += sz + sz {
			s.merge(a, lo, lo+sz-1, min(lo+sz+sz-1, length-1))
		}
	}
}

func Test_MergeSort(t *testing.T) {
	a := []int{3, 7, 2, 9, 10, 1, 6, 5, 4, 8}
	s := &MergeSort{}
	//s.Sort(a)
	s.ButtonUpMergeSort(a)
	t.Log(a)
}
