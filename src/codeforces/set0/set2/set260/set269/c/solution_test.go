package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := process(reader)

	f := make([]int, n)
	for i, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if res[i] == 0 {
			f[u] += w
			f[v] -= w
		} else {
			f[u] -= w
			f[v] += w
		}
	}
	for i := 1; i < n-1; i++ {
		if f[i] != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
3 2 10
1 2 10
3 1 5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 5
1 2 10
1 3 10
2 3 5
4 2 15
3 4 5`)
}
