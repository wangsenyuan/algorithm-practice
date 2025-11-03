package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, d int, h int, expect bool) {
	res := solve(n, d, h)
	if len(res) == n-1 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	g := NewGraph(n, 2*n)
	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	H := make([]int, n)

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		res := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				H[v] = H[u] + 1
				res += dfs(u, v)
			}
		}
		return res
	}

	sz := dfs(-1, 0)

	if sz != n {
		t.Fatalf("Sample result %v, not a tree", res)
	}

	if slices.Max(H) != h {
		t.Fatalf("Sample expect %d, but got %v", h, H)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 16, 15, 14, true)
}
