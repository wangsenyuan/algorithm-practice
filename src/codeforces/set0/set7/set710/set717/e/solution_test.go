package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	color, edges, res := process(reader)

	n := len(color)
	g := make([]map[int]struct{}, n+1)
	for i := range n + 1 {
		g[i] = make(map[int]struct{})
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		g[u][v] = struct{}{}
		g[v][u] = struct{}{}
	}

	x := slices.Min(color)
	if x == 1 {
		if len(res) == 1 {
			return
		}
		t.Fatalf("when all black, expect 1 only, but got %v", res)
	}

	if res[0] != 1 {
		t.Fatalf("expect 1, but got %d", res[0])
	}

	cur := 1
	for i := 1; i < len(res); i++ {
		v := res[i]
		if _, ok := g[cur][v]; !ok {
			t.Fatalf("Sample result %v is not a valid walk at %d", res, i)
		}
		color[v-1] *= -1
		cur = v
	}

	x = slices.Min(color)
	if x != 1 {
		t.Fatalf("Sample result %v, not all black at the end, min color %d", res, x)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1
1
-1
1
-1
2 5
4 3
2 4
4 1
`
	runSample(t, s)
}
