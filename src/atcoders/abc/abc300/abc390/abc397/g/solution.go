package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, k, edges)
}

func solve(n int, k int, edges [][]int) int {
	// A simple path has at most n-1 edges, so distance n is impossible.
	lo, hi := 0, n
	for lo+1 < hi {
		mid := (lo + hi) / 2
		if minMarks(n, edges, mid) <= k {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func minMarks(n int, edges [][]int, dist int) int {
	if dist == 0 {
		return 0
	}

	// node(v, j) represents the Boolean condition level[v] >= j+1.
	// Being on the source side of the cut means the condition is true.
	node := func(v int, j int) int {
		return v*dist + j
	}

	source := n * dist
	sink := source + 1
	flow := NewDinic(sink + 1)
	inf := len(edges) + 1

	// The original source has level 0, and the destination has level dist.
	for j := range dist {
		flow.addEdge(node(0, j), sink, inf)
		flow.addEdge(source, node(n-1, j), inf)
	}

	// level[v] >= j+2 implies level[v] >= j+1.
	for v := range n {
		for j := 0; j+1 < dist; j++ {
			flow.addEdge(node(v, j+1), node(v, j), inf)
		}
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1

		// If level[v] = level[u]+1, exactly one of these edges crosses
		// the cut, representing that the original edge must be marked.
		for j := range dist {
			flow.addEdge(node(v, j), node(u, j), 1)
		}

		// One original edge cannot increase the level by two or more.
		for j := 0; j+1 < dist; j++ {
			flow.addEdge(node(v, j+1), node(u, j), inf)
		}
	}

	return flow.maxFlow(source, sink)
}

type FlowEdge struct {
	to  int
	rev int
	cap int
}

type Dinic struct {
	g     [][]FlowEdge
	level []int
	iter  []int
}

func NewDinic(n int) *Dinic {
	return &Dinic{
		g:     make([][]FlowEdge, n),
		level: make([]int, n),
		iter:  make([]int, n),
	}
}

func (d *Dinic) addEdge(from int, to int, cap int) {
	fwd := FlowEdge{to, len(d.g[to]), cap}
	rev := FlowEdge{from, len(d.g[from]), 0}
	d.g[from] = append(d.g[from], fwd)
	d.g[to] = append(d.g[to], rev)
}

func (d *Dinic) bfs(source int, sink int) bool {
	for i := range d.level {
		d.level[i] = -1
	}
	d.level[source] = 0
	queue := make([]int, 1, len(d.g))
	queue[0] = source
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		for _, e := range d.g[u] {
			if e.cap > 0 && d.level[e.to] < 0 {
				d.level[e.to] = d.level[u] + 1
				queue = append(queue, e.to)
			}
		}
	}
	return d.level[sink] >= 0
}

func (d *Dinic) dfs(u int, sink int, pushed int) int {
	if u == sink {
		return pushed
	}
	for d.iter[u] < len(d.g[u]) {
		i := d.iter[u]
		e := &d.g[u][i]
		if e.cap > 0 && d.level[e.to] == d.level[u]+1 {
			flow := d.dfs(e.to, sink, min(pushed, e.cap))
			if flow > 0 {
				e.cap -= flow
				d.g[e.to][e.rev].cap += flow
				return flow
			}
		}
		d.iter[u]++
	}
	return 0
}

func (d *Dinic) maxFlow(source int, sink int) int {
	var result int
	inf := int(^uint(0) >> 1)
	for d.bfs(source, sink) {
		clear(d.iter)
		for {
			flow := d.dfs(source, sink, inf)
			if flow == 0 {
				break
			}
			result += flow
		}
	}
	return result
}
