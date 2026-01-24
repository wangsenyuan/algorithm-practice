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

	adj := make([][]int, n)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		c := int(res[u] - '0')
		if c == w {
			adj[u] = append(adj[u], v)
		}
	}

	dp := make([]int, n)
	for i := range n {
		dp[i] = -1
	}
	dp[0] = 0
	var que []int
	que = append(que, 0)
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, v := range adj[u] {
			if dp[v] == -1 {
				dp[v] = dp[u] + 1
				que = append(que, v)
			}
		}
	}
	if dp[n-1] != best {
		t.Fatalf("Sample result %s, not correct, %d != %d", res, dp[n-1], best)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 2 0
1 3 1
2 3 0
2 3 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 8
1 1 0
1 3 0
1 3 1
3 2 0
2 1 0
3 4 1
2 4 0
2 4 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 10
1 2 0
1 3 1
1 4 0
2 3 0
2 3 1
2 5 0
3 4 0
3 4 1
4 2 1
4 5 0
`
	expect := -1
	runSample(t, s, expect)
}
