package Sort

import (
	"testing"
)

func Test_SelectSort(t *testing.T) {
	a := []int{9, 4, 1, 6, 0, 7, 8, 5, 3, 2}
	t.Log(0, "=>", a)
	min := func(a, b int) bool { return a < b }

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if min(a[j], a[i]) {
				a[j], a[i] = a[i], a[j]
			}
		}
		t.Log(i, "=>", a)
	}
	t.Log("Result:", a)
}
