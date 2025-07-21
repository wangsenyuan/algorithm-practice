package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, edges, color := process(reader)
	yes_no := readString(reader)

	if (yes_no == "YES") != (res > 0) {
		t.Fatalf("Sample expect %s, but got %d", yes_no, res)
	}
	if yes_no == "NO" {
		return
	}

	n := len(color)

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var dfs func(p int, u int) bool
	dfs = func(p int, u int) bool {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if !dfs(u, v) || color[u] != color[v] {
					return false
				}
			}
		}
		return true
	}

	res--

	for i := g.nodes[res]; i > 0; i = g.next[i] {
		v := g.to[i]
		if !dfs(res, v) {
			t.Fatalf("Sample result %d, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1 2
2 3
3 4
1 2 1 1
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2
2 3
1 2 3
YES`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
1 2
2 3
3 4
1 2 1 2
NO`)
}
