package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, edges := process(reader)
	expect := readString(reader)

	if res == expect {
		return
	}
	if expect == "Impossible!" || res == "Impossible!" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	n := len(edges) + 1

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n)

	fa := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				fa[v] = u
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}
	dfs(0, 0)

	check := func(u int, v int) bool {
		c := res[2*u]
		for u != v {
			if res[2*u] < c || res[2*v] < c {
				return true
			}

			if dep[u] < dep[v] {
				u, v = v, u
			}
			u = fa[u]
		}

		return res[2*u] < c
	}

	for u := range n {
		for v := u + 1; v < n; v++ {
			if res[2*u] == res[2*v] && !check(u, v) {
				t.Fatalf("Sample result %s, not correct", res)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	runSample(t, `10
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
D C B A D C B D C D
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 2
1 3
1 4
A B B B
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
1 2
2 4
4 5
6 4
3 2
B A B B C C
	`)
}
