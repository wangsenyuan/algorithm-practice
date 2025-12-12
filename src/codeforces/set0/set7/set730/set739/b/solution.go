package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	edges := make([][]int, n-1)
	for i := 1; i < n; i++ {
		edges[i-1] = make([]int, 2)
		fmt.Fscan(reader, &edges[i-1][0], &edges[i-1][1])
	}
	return solve(a, edges)
}

func solve(a []int, edges [][]int) []int {
	n := len(a)

	g := NewGraph(n, n)
	for i := 1; i < n; i++ {
		p, w := edges[i-1][0]-1, edges[i-1][1]
		g.AddEdge(p, i, w)
	}

	level := make([]int, n)
	dist := make([]int, n)
	diff := make([]int, n)

	find := func(k int, w int) int {
		return sort.Search(k, func(i int) bool {
			return dist[level[i]] >= w
		})
	}

	var dfs func(u int, d int)
	dfs = func(u int, d int) {
		level[d] = u
		// dist[u] - dist[v] <= a[u]
		j := find(d+1, dist[u]-a[u])
		if j != d && d > 0 {
			diff[level[d-1]]++
			if j > 0 {
				diff[level[j-1]]--
			}
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			dist[v] = dist[u] + w
			dfs(v, d+1)
			diff[u] += diff[v]
		}
	}

	dfs(0, 0)

	return diff
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
