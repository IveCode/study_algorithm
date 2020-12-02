package Sort

import "testing"

func Test_ShellSort(t *testing.T) {
	a := []int{9, 4, 1, 6, 0, 7, 8, 5, 3, 2}
	//h := len(a)/2 + 1
	h := 1
	for h <= len(a) {
		h = h*3 + 1
	}
	min := func(a, b int) bool { return a < b }
	for h >= 1 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h; j -= h {
				if min(a[j], a[j-h]) {
					a[j], a[j-h] = a[j-h], a[j]
				}
			}
			t.Log(h, "=>", a)
		}
		//h = h / 2
		h = h / 3
	}
	t.Log(a)
}
