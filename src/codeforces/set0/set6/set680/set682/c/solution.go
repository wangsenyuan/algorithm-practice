package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		for j := range 2 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	return solve(a, edges)
}

func solve(a []int, edges [][]int) int {
	n := len(a)
	g := NewGraph(n, n)

	for i, e := range edges {
		p, w := e[0]-1, e[1]
		g.AddEdge(p, i+1, w)
	}

	sz := make([]int, n)
	var dfs func(u int)
	dfs = func(u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
	}

	dfs(0)

	var res int
	var dfs2 func(u int, up int, dist int)
	dfs2 = func(u int, up int, dist int) {
		if dist-a[u] > up {
			res += sz[u]
			return
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			dfs2(v, min(up, dist+w), dist+w)
		}
	}

	dfs2(0, 0, 0)

	return res
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type Item struct {
	id       int
	priority int
	index    int
}
