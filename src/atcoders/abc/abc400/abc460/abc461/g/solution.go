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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := NewFlowGraph(2*n+2, 0, 2*n+1)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, n+v, 1)
		g.AddEdge(v, n+u, 1)
	}

	for i := 1; i <= n; i++ {
		g.AddEdge(0, i, 1)
		g.AddEdge(n+i, 2*n+1, 1)
	}

	return n*2026 - g.MaxFlow()*1013
}

const INF = 1 << 60

type FlowGraph struct {
	n            int
	edges        []*Edge
	adj          [][]int
	que          []int
	dist         []int
	source, sink int
}

func NewFlowGraph(n int, source int, sink int) FlowGraph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 10)
	}
	edges := make([]*Edge, 0, n)
	return FlowGraph{n, edges,
		adj, make([]int, n), make([]int, n), source, sink}
}

func (g *FlowGraph) AddEdge(u, v, c int) {
	g.adj[u] = append(g.adj[u], len(g.edges))
	e1 := &Edge{u, v, c, 0}
	g.edges = append(g.edges, e1)

	g.adj[v] = append(g.adj[v], len(g.edges))
	e2 := &Edge{v, u, 0, 0}
	g.edges = append(g.edges, e2)
}

func (g FlowGraph) bfs() bool {
	for i := 0; i < g.n; i++ {
		g.dist[i] = -1
	}
	var head, tail int

	g.que[tail] = g.source
	tail++
	g.dist[g.source] = 0

	for head < tail {
		u := g.que[head]
		head++
		for _, id := range g.adj[u] {
			e := g.edges[id]
			if g.dist[e.to] < 0 && e.flow < e.capacity {
				g.que[tail] = e.to
				tail++
				g.dist[e.to] = g.dist[u] + 1
			}
		}
	}
	return g.dist[g.sink] != -1
}

func (g *FlowGraph) dfs(v int, flow int, ptr []int) int {
	if flow == 0 {
		return 0
	}
	if v == g.sink {
		return flow
	}

	for ptr[v] < len(g.adj[v]) {
		id := g.adj[v][ptr[v]]
		ptr[v]++
		edge := g.edges[id]
		if g.dist[edge.to] != g.dist[edge.from]+1 {
			continue
		}
		pushed := g.dfs(edge.to, min(flow, edge.capacity-edge.flow), ptr)

		if pushed > 0 {
			g.edges[id].flow += pushed
			g.edges[id^1].flow -= pushed
			return pushed
		}
	}
	return 0
}

func (g *FlowGraph) MaxFlow() int {
	var flow int
	ptr := make([]int, g.n)
	for g.bfs() {
		for i := 0; i < g.n; i++ {
			ptr[i] = 0
		}
		for {
			pushed := g.dfs(g.source, INF, ptr)
			if pushed == 0 {
				break
			}
			flow += pushed
		}
	}

	return flow
}

type Edge struct {
	from, to       int
	capacity, flow int
}
