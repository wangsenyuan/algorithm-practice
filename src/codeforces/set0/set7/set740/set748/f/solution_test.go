package main

import (
	"bufio"
	"math/bits"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, m int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, c, d, res := drive(reader)

	if len(d) != m {
		t.Fatalf("Expected %d centers, but got %d", m, len(d))
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dep := make([]int, n)
	h := bits.Len(uint(n)) + 1
	fa := make([][]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[fa[u][i]] >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
			}
		}
		return fa[u][0]
	}

	dist := func(u int, v int) int {
		return dep[u] + dep[v] - 2*dep[lca(u, v)]
	}

	marked := make([]bool, n)
	for _, v := range c {
		marked[v-1] = true
	}

	for _, cur := range res {
		u, v, x := cur[0]-1, cur[1]-1, cur[2]-1
		if !marked[u] || !marked[v] {
			t.Fatalf("%d or %d is not the tour city", u, v)
		}

		d1 := dist(u, v)
		d2 := dist(u, x) + dist(x, v)

		if d1 != d2 {
			t.Fatalf("%v not a valid answer", cur)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 2
1 2
1 3
2 4
2 5
3 6
2 5 4 6`
	m := 1

	runSample(t, s, m)
}

func TestSample2(t *testing.T) {
	s := `6 2
1 6
6 2
6 5
5 3
5 4
1 3 4 2`
	m := 1

	runSample(t, s, m)
}
