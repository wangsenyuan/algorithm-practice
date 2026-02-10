package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, edges, a, ok, res := drive(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}

	adj := make([]map[int]bool, n)
	for i := range n {
		adj[i] = make(map[int]bool)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = true
		adj[v][u] = true
	}
	occ := make([]int, n)
	if len(res) > 0 {
		occ[res[0]-1] = 1
		for i := 0; i+1 < len(res); i++ {
			u, v := res[i]-1, res[i+1]-1
			if !adj[u][v] {
				t.Fatalf("Sample result %v, not valid", res)
			}
			occ[v] ^= 1
		}
	}
	if !slices.Equal(occ, a) {
		t.Fatalf("Sample result %v, not valid", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2
2 3
1 1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 7
1 2
1 3
1 4
1 5
3 4
3 5
4 5
0 1 0 1 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 0
0 0
`
	expect := true
	runSample(t, s, expect)
}
