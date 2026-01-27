package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	best, res := solve(n)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
	ans := res[0]
	for i := 1; i < n; i++ {
		if i&1 == 1 {
			ans = ans | res[i]
		} else {
			ans = ans & res[i]
		}
	}
	if ans != best {
		t.Fatalf("Sample %d result %v, not correct", n, res)
	}
}

func TestSample(t *testing.T) {
	// 1 | 2 & 3 | 4
	cases := [][]int{
		{4, 7},
		{5, 5},
		{6, 7},
		{7, 7},
		{8, 15},
		{9, 9},
		{10, 15},
	}
	for _, cur := range cases {
		runSample(t, cur[0], cur[1])
	}
}
