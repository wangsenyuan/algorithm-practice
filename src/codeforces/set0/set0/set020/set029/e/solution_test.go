package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, k, paths := drive(reader)
	if k != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, k)
	}
	if expect < 0 {
		return
	}

	if paths[0][0] != 1 || paths[1][0] != n {
		t.Fatalf("Sample result %v, not correct", paths)
	}

	if paths[0][k] != n || paths[1][k] != 1 {
		t.Fatalf("Sample result %v, not correct", paths)
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

	for i := range k {
		u, v := paths[0][i]-1, paths[1][i]-1
		u1, v1 := paths[0][i+1]-1, paths[1][i+1]-1
		if !adj[u][u1] || !adj[v][v1] {
			t.Fatalf("Sample result %v, not correct", paths)
		}
		if u1 == v1 {
			t.Fatalf("Sample result %v, not correct", paths)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `2 1
1 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 5
1 2
2 7
7 6
2 3
3 4
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 6
1 2
2 7
7 6
2 3
3 4
1 5
`
	expect := 6
	runSample(t, s, expect)
}

func TestDenseGraphWithoutDirectEdge(t *testing.T) {
	var buf strings.Builder
	n := 140
	m := n*(n-1)/2 - 1
	fmt.Fprintf(&buf, "%d %d\n", n, m)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if i == 1 && j == n {
				continue
			}
			fmt.Fprintf(&buf, "%d %d\n", i, j)
		}
	}
	runSample(t, buf.String(), 2)
}
