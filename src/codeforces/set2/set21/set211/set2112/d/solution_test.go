package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, _, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	// 有向边组成
	adj := make([][]int, n)
	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
	}

	dp := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		if dp[u] > 0 {
			return
		}
		dp[u] = 1
		for _, v := range adj[u] {
			dfs(v)
			dp[u] += dp[v]
		}
	}

	for u := range n {
		dfs(u)
	}

	var sum int

	for u := range n {
		// 以u为起点的简单路径的数量
		sum += dp[u] - 1
	}

	if sum != n {
		t.Fatalf("Sample expect %d, but got %d", n, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
2 4
1 3
3 5`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
1 3
1 4
4 5`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
3 1
1 2
2 4`
	expect := true
	runSample(t, s, expect)
}
