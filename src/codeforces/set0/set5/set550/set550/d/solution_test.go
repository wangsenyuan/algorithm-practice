package main

import "testing"

func runSample(t *testing.T, k int, expect bool) {
	n, res := solve(k)

	if n > 0 != expect {
		t.Fatalf("n = %d, expect %t", n, expect)
	}
	if !expect {
		return
	}

	deg := make([]int, n+1)

	for _, edge := range res {
		u, v := edge[0], edge[1]
		deg[u]++
		deg[v]++
	}

	for i := 1; i <= n; i++ {
		if deg[i] != k {
			t.Fatalf("deg[%d] = %d, want %d", i, deg[i], k)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, true)
}
