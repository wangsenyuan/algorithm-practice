package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := drive(reader)

	if len(res) == n != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}
	adj := make([][]int, n)
	for i := range n {
		adj[i] = make([]int, n)
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = 1
		adj[v][u] = 1
	}

	check := func(x byte, y byte) int {
		if x == 'b' || y == 'b' || x == y {
			return 1
		}
		return 0
	}

	for i := range n {
		for j := range i {
			w := check(res[i], res[j])
			if w != adj[j][i] {
				t.Fatalf("Sample result %s, not valid", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
1 2
1 3
1 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
1 2
1 3
1 4
3 4
`
	expect := true
	runSample(t, s, expect)
}
