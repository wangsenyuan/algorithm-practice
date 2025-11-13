package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, roads, res := drive(reader)

	g := NewGraph(n, len(roads))

	for _, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v)
	}

	dist := make([]int, n)

	que := make([]int, n)
	bfs := func(s int) {
		for i := range n {
			dist[i] = -1
		}
		dist[s] = 0
		var head, tail int
		que[head] = s
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
	}

	check := func(arr []int) int {
		var res int
		for i := range 3 {
			bfs(arr[i] - 1)
			nxt := arr[i+1] - 1
			if dist[nxt] == -1 {
				t.Fatalf("Sample result %v, not a valid path", arr)
			}
			res += dist[nxt]
		}
		return res
	}

	x := check(expect)
	y := check(res)
	if x != y {
		t.Fatalf("Sample result %v, not correct, expect %d, but got %d", res, x, y)
	}
}

func TestSample1(t *testing.T) {
	s := `8 9
1 2
2 3
3 4
4 1
4 5
5 6
6 7
7 8
8 5
`
	expect := []int{2, 1, 8, 7}
	runSample(t, s, expect)
}
