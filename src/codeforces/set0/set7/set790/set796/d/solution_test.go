package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, d, edges, stations, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	marked := make([]bool, n-1)
	for _, v := range res {
		marked[v-1] = true
	}

	g := NewGraph(n, 2*n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		if !marked[i] {
			g.AddEdge(u, v, i)
			g.AddEdge(v, u, i)
		}
	}

	dist := make([]int, n)
	for i := range n {
		dist[i] = -1
	}
	que := make([]int, n)

	var head, tail int
	for _, s := range stations {
		s--
		if dist[s] < 0 {
			dist[s] = 0
			que[head] = s
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}

	for i := range n {
		if dist[i] < 0 || dist[i] > d {
			t.Fatalf("Sample result not correct, it cause %d to have distance %d from station more than %d", i, dist[i], d)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 2 4
1 6
1 2
2 3
3 4
4 5
5 6
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3 2
1 5 6
1 2
1 3
1 4
1 5
5 6
`
	expect := 2
	runSample(t, s, expect)
}
