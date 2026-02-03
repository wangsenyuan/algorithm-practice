package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, k, a, res := drive(reader)
	if len(res) == n-1 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	// n > 1

	// 需要计算叶子节点和d
	adj := make([][]int, n)
	for _, cur := range res {
		u, v := cur[0]-1, cur[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	cnt := make([]int, len(a))

	var leaf int
	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		if len(adj[u]) == 1 && p >= 0 {
			leaf++
		}
		if d > 0 {
			cnt[d-1]++
		}
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v, d+1)
			}
		}
	}

	dfs(-1, 0, 0)

	if leaf != k || !slices.Equal(cnt, a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 3 3
2 3 1`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `14 5 6
4 4 2 2 1`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1 1
2`
	expect := false
	runSample(t, s, expect)
}
