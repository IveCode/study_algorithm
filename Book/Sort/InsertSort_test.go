package Sort

import "testing"

func Test_InsertSort(t *testing.T) {
	a := []int{9, 4, 1, 6, 0, 7, 8, 5, 3, 2}
	t.Log(0, "=>", a)
	min := func(a, b int) bool { return a < b }
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if min(a[j], a[j-1]) {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
		t.Log(i, "=>", a)
	}
	t.Log("Result:", a)
}
