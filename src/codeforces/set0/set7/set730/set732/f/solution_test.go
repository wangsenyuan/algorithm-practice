package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, _, best, res := drive(reader)
	if int(best) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	adj := make([][]int, n)
	adj2 := make([][]int, n)

	add := func(u int, v int) {
		adj[u] = append(adj[u], v)
		adj2[v] = append(adj2[v], u)
	}

	for _, e := range res {
		u, v := int(e[0]-1), int(e[1]-1)
		add(u, v)
	}

	var order []int

	var dfs func(u int)

	vis := make([]bool, n)

	dfs = func(u int) {
		vis[u] = true
		for _, v := range adj[u] {
			if !vis[v] {
				dfs(v)
			}
		}
		order = append(order, u)
	}

	for u := range int(n) {
		if !vis[u] {
			dfs(u)
		}
	}

	clear(vis)

	comp := make([]int, n)
	var cnt []int

	var dfs2 func(u int)
	dfs2 = func(u int) {
		if vis[u] {
			return
		}
		comp[u] = len(cnt) - 1
		cnt[len(cnt)-1]++
		vis[u] = true
		for _, v := range adj2[u] {
			dfs2(v)
		}
	}

	for i := int(n) - 1; i >= 0; i-- {
		u := order[i]
		if !vis[u] {
			cnt = append(cnt, 0)
			dfs2(u)
		}
	}

	topo := make([][]int, len(cnt))

	for _, e := range res {
		u, v := int(e[0]-1), int(e[1]-1)
		if comp[u] != comp[v] {
			topo[comp[u]] = append(topo[comp[u]], comp[v])
		}
	}

	dp := make([]int, len(cnt))

	var dfs3 func(u int) int
	dfs3 = func(u int) int {
		if dp[u] > 0 {
			return dp[u]
		}
		dp[u] = cnt[u]
		for _, v := range topo[u] {
			dp[u] += dfs3(v)
		}
		return dp[u]
	}

	for i := range len(cnt) {
		dp[i] = dfs3(i)
	}

	if slices.Min(dp) < int(best) {
		t.Fatalf("Sample result %v, not correct, %d <  %d, not best", res, slices.Min(dp), best)
	}
}

func TestSample1(t *testing.T) {
	s := `7 9
4 3
2 6
7 1
4 1
7 3
3 5
7 4
6 5
2 5
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
2 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12 16
12 3
10 12
12 9
4 10
1 12
10 6
2 4
7 10
3 8
9 8
10 5
1 11
2 11
11 10
12 2
3 10
`
	expect := 9
	runSample(t, s, expect)
}
