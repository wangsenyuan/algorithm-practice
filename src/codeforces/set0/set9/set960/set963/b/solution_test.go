package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, p, res := process(reader)

	expect := readString(reader)

	if len(res) > 0 != (expect == "YES") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}

	if expect != "YES" {
		return
	}

	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		if p[i] == 0 {
			continue
		}
		g.AddEdge(p[i]-1, i)
		g.AddEdge(i, p[i]-1)
		deg[p[i]-1]++
		deg[i]++
	}

	for _, u := range res {
		u--
		if deg[u]%2 == 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
0 1 2 1 2
YES
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
0 1 2 3
NO
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
4 4 2 0 2
YES
`)
}
