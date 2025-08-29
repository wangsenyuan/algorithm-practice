package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, loan, res := drive(reader)
	g := make([]map[int]bool, n+1)
	for i := range n + 1 {
		g[i] = make(map[int]bool)
	}
	for _, cur := range loan {
		g[cur[0]][cur[1]] = true
	}

	if len(res) != n {
		t.Fatalf("Sample result %v, not correct, len(res) = %d, want %d", res, len(res), n)
	}

	for i := 0; i+1 < len(res); i++ {
		u, v := res[i], res[i+1]
		if g[u][v] {
			t.Fatalf("Sample result %v, not correct, invite %d and %d in sequence, will make %d unhappy", res, u, v, u)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
1 2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2
2 3
3 1
`
	runSample(t, s)
}
