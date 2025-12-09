package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, best, res := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	m := len(edges)
	rev := make([]bool, m)
	for _, i := range res {
		rev[i-1] = true
	}

	g := NewGraph(n, m)
	for i, e := range edges {
		u, v, c := e[0]-1, e[1]-1, e[2]
		if rev[i] {
			if c > best {
				t.Fatalf("Sample result %v, not allowed to reverse edge %d", res, i+1)
			}
			g.AddEdge(v, u, i)
		} else {
			g.AddEdge(u, v, i)
		}
	}

	marked := make([]int, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		marked[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if marked[v] == 1 || marked[v] == 0 && !dfs(v) {
				return false
			}
		}
		marked[u]++
		return true
	}

	for i := range n {
		if marked[i] == 0 && !dfs(i) {
			t.Fatalf("Sample result %v, found a cycle from %d", res, i+1)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `5 6
2 1 1
5 2 6
2 3 2
3 4 3
4 5 5
1 5 4
`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `5 7
2 1 5
3 2 3
1 3 3
2 4 1
4 3 5
5 4 1
1 5 3
`
	runSample(t, s, 3)
}

func TestSample3(t *testing.T) {
	s := `10 45
5 6 8
9 4 8
7 1 1
9 7 1
7 2 1
1 4 2
5 7 7
10 5 7
7 8 8
8 5 4
4 7 3
1 8 7
3 1 9
9 1 3
10 2 7
6 2 7
2 5 7
5 4 7
6 7 6
4 2 7
6 8 10
6 10 2
3 6 3
10 3 6
4 3 6
3 9 8
5 1 4
2 3 7
3 8 1
9 10 4
9 8 7
4 6 6
2 8 1
7 3 5
9 5 4
7 10 2
4 8 8
10 4 10
10 8 5
10 1 10
5 3 9
9 6 2
6 1 5
2 1 1
2 9 5
`
	runSample(t, s, 7)
}

func TestSample4(t *testing.T) {
	s := `1000 10
868 438 2
343 550 7
398 889 5
124 36 2
135 199 5
457 601 3
399 457 5
207 830 1
993 9 6
94 532 2
`
	runSample(t, s, 0)
}