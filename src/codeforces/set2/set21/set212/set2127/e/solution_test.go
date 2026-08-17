package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectCost int64) {
	t.Helper()
	k, w, orig, edges := parseCase(s)
	cost, colors := drive(bufio.NewReader(strings.NewReader(s)))
	if !validColors(k, orig, colors) {
		t.Fatalf("invalid coloring %v for original %v and k=%d", colors, orig, k)
	}
	got := cutieCost(w, colors, edges)
	if got != cost {
		t.Fatalf("reported cost %d, but coloring %v has cutie cost %d", cost, colors, got)
	}
	if cost != expectCost {
		t.Fatalf("Sample expect cost %d, but got %d (colors %v)", expectCost, cost, colors)
	}
}

func parseCase(s string) (k int, w, c []int, edges [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))
	var n int
	fmt.Fscan(reader, &n, &k)
	w = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &w[i])
	}
	c = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return
}

func validColors(k int, orig, got []int) bool {
	if len(got) != len(orig) {
		return false
	}
	for i, col := range got {
		if orig[i] != 0 {
			if col != orig[i] {
				return false
			}
			continue
		}
		if col < 1 || col > k {
			return false
		}
	}
	return true
}

func cutieCost(w, colors []int, edges [][]int) int64 {
	n := len(w)
	parent, depth := rootedTree(n, edges)
	lca := func(u, v int) int {
		for depth[u] > depth[v] {
			u = parent[u]
		}
		for depth[v] > depth[u] {
			v = parent[v]
		}
		for u != v {
			u = parent[u]
			v = parent[v]
		}
		return u
	}
	cutie := make([]bool, n)
	for x := range n {
		for y := x + 1; y < n; y++ {
			if colors[x] != colors[y] {
				continue
			}
			v := lca(x, y)
			if colors[x] != colors[v] {
				cutie[v] = true
			}
		}
	}
	var cost int64
	for v, ok := range cutie {
		if ok {
			cost += int64(w[v])
		}
	}
	return cost
}

func rootedTree(n int, edges [][]int) (parent, depth []int) {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	parent = make([]int, n)
	depth = make([]int, n)
	var dfs func(u, p int)
	dfs = func(u, p int) {
		parent[u] = p
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			depth[v] = depth[u] + 1
			dfs(v, u)
		}
	}
	dfs(0, -1)
	return parent, depth
}

func TestSample1(t *testing.T) {
	runSample(t, `4 4
5 5 5 5
1 0 2 3
1 2
1 3
1 4
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
3 1 4 1 5
1 2 1 2 2
1 4
2 1
3 4
4 5
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `11 3
3 1 4 3 1 4 3 1 4 5 6
0 0 0 2 1 2 1 2 2 1 1
1 2
2 3
2 4
2 5
2 6
1 7
7 8
7 9
7 10
7 11
`, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 3
2 3 2 3
2 1 0 0
3 1
1 2
2 4
`, 0)
}

func TestCutieCostNoteColorings(t *testing.T) {
	s := `11 3
3 1 4 3 1 4 3 1 4 5 6
0 0 0 2 1 2 1 2 2 1 1
1 2
2 3
2 4
2 5
2 6
1 7
7 8
7 9
7 10
7 11
`
	_, w, _, edges := parseCase(s)
	cases := []struct {
		colors []int
		cost   int64
	}{
		{[]int{3, 1, 2, 2, 1, 2, 1, 2, 2, 1, 1}, 7},
		{[]int{3, 2, 1, 2, 1, 2, 1, 2, 2, 1, 1}, 7},
		{[]int{2, 3, 1, 2, 1, 2, 1, 2, 2, 1, 1}, 7},
		{[]int{2, 2, 2, 2, 1, 2, 1, 2, 2, 1, 1}, 6},
	}
	for _, tc := range cases {
		if got := cutieCost(w, tc.colors, edges); got != tc.cost {
			t.Fatalf("coloring %v: cutie cost %d, want %d", tc.colors, got, tc.cost)
		}
	}
}
