package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	nodes, res := drive(reader)

	n := len(nodes)

	if len(res) != n-1 {
		t.Fatalf("Sample result %v, not a valid tree", res)
	}

	set := NewDSU(n)

	sum := make([]int, n)

	for _, cur := range res {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		if !set.Union(u, v) {
			t.Fatalf("Sample result %v, having cycle", res)
		}
		if nodes[u][0] == nodes[v][0] {
			t.Fatalf("Sample result %v, connecting same color node", res)
		}
		sum[u] += w
		sum[v] += w
	}

	for i := range n {
		if sum[i] != nodes[i][1] {
			t.Fatalf("Sample reuslt %v, not valid", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3
1 2
0 5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6
1 0
0 3
1 8
0 2
0 3
0 0
`
	runSample(t, s)
}
