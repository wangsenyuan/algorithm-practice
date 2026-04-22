package main

import "testing"

func runSample(t *testing.T, l []int, expect int) {
	n := len(l)

	var cnt int

	ask := func(w int) int {
		if cnt > n+30 {
			t.Fatalf("Sample asked too much times %d", cnt)
		}
		var h int
		for i := 0; i < n; {
			var cur int
			for i < n && cur+l[i] <= w {
				cur += l[i]
				i++
				if cur == w {
					break
				}
				// 再增加一个space
				cur++
			}
			if cur == 0 {
				// 无法inc i, w < l[i]
				return 0
			}
			h++
		}
		cnt++
		return h
	}

	res := solve(n, ask)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := []int{5, 2, 7, 3, 5, 6}
	expect := 32
	runSample(t, l, expect)
}
