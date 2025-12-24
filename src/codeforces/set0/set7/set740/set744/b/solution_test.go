package main

import "testing"

func runSample(t *testing.T, M [][]int) {
	n := len(M)
	var cnt int
	ask := func(w []int) []int {
		cnt++
		if cnt > 20 {
			t.Fatalf("Sample asked too much")
		}
		res := make([]int, n)
		for i := range n {
			res[i] = inf
			for _, j := range w {
				res[i] = min(res[i], M[i][j-1])
			}
		}
		return res
	}

	ans := solve(n, ask)

	for i := range n {
		res := inf
		for j := range n {
			if i != j {
				res = min(res, M[i][j])
			}
		}
		if res != ans[i] {
			t.Fatalf("Sample expect %v, but got %v", ans, res)
		}
	}
}

func TestSample1(t *testing.T) {
	M := [][]int{
		{0, 3, 2},
		{5, 0, 7},
		{4, 8, 0},
	}
	runSample(t, M)
}
