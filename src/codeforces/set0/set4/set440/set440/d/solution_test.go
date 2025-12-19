package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, k, edges, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	marked := make([]bool, len(edges))
	for _, v := range res {
		marked[v-1] = true
	}

	adj := make([][]int, n)
	for i, e := range edges {
		if !marked[i] {
			u, v := e[0]-1, e[1]-1
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
	}

	vis := make([]bool, n)

	var dfs func(u int) int
	dfs = func(u int) int {
		res := 1
		vis[u] = true
		for _, v := range adj[u] {
			if !vis[v] {
				res += dfs(v)
			}
		}
		return res
	}

	ok := false
	for u := range n {
		if !vis[u] {
			sz := dfs(u)
			if sz == k {
				ok = true
			}
		}
	}

	if !ok {
		t.Fatalf("Sample result %v, not correct, no size(%d) state found", res, k)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
1 2
2 3
3 4
4 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
1 2
1 3
1 4
1 5
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `11 4
1 2
1 3
1 4
2 6
2 7
1 5
2 8
4 9
4 10
4 11`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `15 5
9 12
8 9
12 14
13 9
15 8
10 9
9 4
1 9
1 3
3 2
9 11
15 6
1 5
15 7`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `17 7
3 13
13 15
15 2
11 13
3 16
8 15
17 11
1 3
9 16
15 10
8 6
12 8
5 1
15 4
11 7
14 6`
	expect := 2
	runSample(t, s, expect)
}
