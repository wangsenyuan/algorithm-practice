package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := drive(reader)

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dist := make([]int, n)

	in := make([]int, n)
	var timer int
	sz := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		in[u] = timer
		timer++
		sz[u] = 1
		for _, v := range adj[u] {
			if p != v {
				dist[v] = dist[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, n-1)

	isAnc := func(u int, v int) bool {
		return in[u] <= in[v] && in[v] < in[u]+sz[u]
	}

	marked := make([]bool, n)

	var cat int

	for i, cur := range res {
		if i > 0 && cur[0] == 2 && res[i-1][0] == 2 {
			t.Fatalf("Sample result %v not valid, there are consecutive destroyes", res)
		}
		if cur[0] == 2 {
			u := cur[1] - 1
			marked[u] = true
			if isAnc(u, cat) {
				t.Fatalf("Sample result %v not valid, cat can't reach home", res)
			}
		} else {
			// move
			next := -1
			for _, v := range adj[cat] {
				if !marked[v] {
					if next < v || dist[v] > dist[next] {
						next = v
					}
				}
			}
			if next == -1 {
				t.Fatalf("Sample result %v not valid, cat can't move", res)
			}
			cat = next
			if cat == n-1 {
				break
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
2 3
1 5
5 4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
1 2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6
1 2
1 3
3 4
4 5
4 6`
	runSample(t, s)
}