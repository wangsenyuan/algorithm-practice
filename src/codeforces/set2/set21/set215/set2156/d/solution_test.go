package main

import "testing"

func runSample(t *testing.T, p []int) {
	t.Helper()
	n := len(p)
	var cnt int
	ask := func(i, x int) int {
		cnt++
		if cnt > 2*n {
			t.Fatalf("too many queries: %d > %d", cnt, 2*n)
		}
		if i < 1 || i > n-1 {
			t.Fatalf("invalid index %d", i)
		}
		if p[i-1]&x == 0 {
			return 0
		}
		return 1
	}

	got := solve(n, ask)
	if got != p[n-1] {
		t.Fatalf("expect %d, but got %d", p[n-1], got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 3, 2})
}

func TestAllSmallPermutations(t *testing.T) {
	for n := 2; n <= 8; n++ {
		p := make([]int, n)
		for i := range n {
			p[i] = i + 1
		}
		var dfs func(int)
		dfs = func(i int) {
			if i == n {
				runSample(t, p)
				return
			}
			for j := i; j < n; j++ {
				p[i], p[j] = p[j], p[i]
				dfs(i + 1)
				p[i], p[j] = p[j], p[i]
			}
		}
		dfs(0)
	}
}
