package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var m, k, p int
	fmt.Fscan(reader, &m, &k, &p)
	edges := make([][]int, m-1)
	for i := range m - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(k, p, edges)
}

func solve(k int, p int, edges [][]int) int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var freq []int

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		if len(freq) == d {
			freq = append(freq, 1)
		} else {
			freq[d]++
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, d+1)
			}
		}
	}

	dfs(0, 0, 0)

	m := len(freq)
	suf := make([]int, m+3)
	for i := m - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + freq[i]
	}

	best := min(slices.Max(freq), k)

	var s int
	var t int

	for l, r := 1, 1; r < m; r++ {
		s += freq[r]
		t += r * freq[r]
		for l < r && r*s-t > p {
			s -= freq[l]
			t -= l * freq[l]
			l++
		}
		tmp := suf[l] - suf[r+1]
		if l > 1 {
			tmp += min(freq[l-1], (p-(r*s-t))/(r-l+1))
		}
		best = max(best, min(tmp, k))
	}

	return best
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
