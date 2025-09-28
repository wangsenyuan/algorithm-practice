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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	m := len(edges)
	g := NewGraph(n, 2*m)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)
	dfn := make([]int, n)
	low := make([]int, n)
	var timer int
	stack := make([]int, n)
	var top int

	id := make([]int, n)
	var scc int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		stack[top] = u
		top++
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if !vis[v] {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else {
				// back edge
				low[u] = min(low[u], dfn[v])
			}
		}
		if low[u] == dfn[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				id[v] = scc
				if v == u {
					break
				}
			}
			scc++
		}
	}

	// 因为原图是连通的，只要从一个节点去处理
	dfs(-1, 0)

	tr := NewGraph(scc, 2*n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		if id[u] != id[v] {
			tr.AddEdge(id[u], id[v])
			tr.AddEdge(id[v], id[u])
		}
	}

	return getDiameter(scc, tr) - 1
}

func getDiameter(n int, g *Graph) int {
	dist := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	far := slices.Max(dist)
	var first int
	for u := range n {
		if dist[u] == far {
			first = u
			break
		}
	}
	dist[first] = 0
	dfs(-1, first)

	return slices.Max(dist) + 1
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
