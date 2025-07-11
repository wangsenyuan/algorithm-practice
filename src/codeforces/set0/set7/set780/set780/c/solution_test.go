package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, k, ans := process(reader)
	expect := readNum(reader)

	if k != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, k)
	}
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]int, k+1)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		vis[ans[u]]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			vis[ans[v]]++
		}

		if vis[ans[u]] > 1 {
			t.Fatalf("Sample result %v, not correct", ans)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[ans[v]] > 1 {
				t.Fatalf("Sample result %v, not correct", ans)
			}
			vis[ans[v]] = 0
		}

		vis[ans[u]] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)
}

func TestSample1(t *testing.T) {
	runSample(t, `3
2 3
1 3
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
2 3
5 3
4 3
1 3
5`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
2 1
3 2
4 3
5 4
3`)
}
