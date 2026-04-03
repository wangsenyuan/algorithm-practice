package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, conditions, res := drive(reader)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if len(expect) == 0 {
		return
	}

	adj := make([][]int, n)
	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}
	dp := make([]int, n)

	var dfs func(p int, u int, sum int)
	dfs = func(p int, u int, sum int) {
		dp[u] = sum
		for _, i := range adj[u] {
			v := (edges[i][0] - 1) ^ (edges[i][1] - 1) ^ u
			if p != v {
				dfs(u, v, sum^res[i])
			}
		}
	}

	dfs(-1, 0, 0)

	for _, cur := range conditions {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		if dp[u]^dp[v] != w {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	var best int
	for _, x := range expect {
		best ^= x
	}
	var got int
	for _, x := range res {
		got ^= x
	}
	if best != got {
		t.Fatalf("Sample result %v, not best %d, got %d", res, best, got)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2
2 3
3 4
1 4 3
2 4 2
1 3 1
2 3 1	`
	runSample(t, s, nil)
}

func TestSample2(t *testing.T) {
	s := `6 2
1 2
2 3
3 4
2 5
5 6
1 4 2
2 6 7
`
	expect := []int{4, 2, 4, 1, 6}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2
1 2
2 3
3 4
2 5
5 6
1 4 3
1 6 5
`
	expect := []int{6, 1, 4, 3, 0}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 5
4 9
1 10
3 6
2 5
8 1
5 9
3 7
6 1
5 3
9 6 326077422
5 9 860153058
3 6 421606914
1 5 35620604
3 8 977423958
`
	expect := []int{0, 0, 421606914, 0, 22516132, 860153058, 806155194, 573893104, 956979470}
	runSample(t, s, expect)
}
