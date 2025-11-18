package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, a, b, edges, ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}

	adj := make([][]bool, n)
	for i := range n {
		adj[i] = make([]bool, n)
		adj[i][i] = true
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = true
		adj[v][u] = true
	}

	c := make([]int, n)
	d := make([]int, n)
	for i := range n {
		for j := range n {
			if res[i][j] > 0 && !adj[i][j] {
				t.Fatalf("Sample result %v, no correct, no roads between %d and %d", res, i, j)
			}
			c[j] += res[i][j]
			d[i] += res[i][j]
		}
	}

	for i := range n {
		if d[i] != a[i] || c[i] != b[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2 6 3
3 5 3 1
1 2
2 3
3 4
4 2
`
	expect := true
	runSample(t, s, expect)
}
