package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, m, x int
	fmt.Fscan(reader, &n, &m, &x)
	edges := make([][]int, m)
	for i := range m {
		var u, v, c int
		fmt.Fscan(reader, &u, &v, &c)
		edges[i] = []int{u, v, c}
	}
	return solve(n, x, edges)
}

func solve(n int, x int, edges [][]int) float64 {

	check := func(w float64) bool {
		if w == 0 {
			return true
		}
		g := NewGraph(n, 2*len(edges))

		addEdge := func(u int, v int, c int) {
			g.AddEdge(u, v, c)
			g.AddEdge(v, u, 0)
		}

		for _, cur := range edges {
			u, v, c := cur[0]-1, cur[1]-1, cur[2]
			y := int(float64(c) / w)
			for w*float64(y+1) <= float64(c) {
				y++
			}
			if y > 0 {
				addEdge(u, v, y)
			}
		}

		return dinic(0, n-1, n, g) >= x
	}

	var l, r float64 = 0, 1e9

	for range 100 {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return float64(x) * (l + r) / 2
}

const inf = 1 << 60

func dinic(src, snk int, n int, g *Graph) int {
	level := make([]int, n)
	all_minus_one := make([]int, n)
	for i := range n {
		all_minus_one[i] = -1
	}

	que := make([]int, n)

	bfs := func() bool {
		var front, end int
		copy(level, all_minus_one)
		level[src] = 0
		que[end] = src
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				if g.limit[i] > g.flow[i] && level[g.to[i]] == -1 {
					v := g.to[i]
					level[v] = level[u] + 1
					que[end] = v
					end++
				}
			}
		}
		return level[snk] > 0
	}

	pos := make([]int, n)

	var dfs func(u int, flow int) int
	dfs = func(u int, flow int) int {
		if flow == 0 {
			return 0
		}
		if u == snk {
			return flow
		}

		for pos[u] > 0 {
			i := pos[u]
			v := g.to[i]
			if level[v] == level[u]+1 && g.flow[i] < g.limit[i] {
				tr := dfs(v, min(flow, g.limit[i]-g.flow[i]))
				if tr > 0 {
					g.flow[i] += tr
					g.flow[i^1] -= tr
					return tr
				}
			}

			pos[u] = g.next[i]
		}
		return 0
	}
	var flow int
	for bfs() {
		for i := 0; i < n; i++ {
			pos[i] = g.node[i]
		}
		for {
			cur := dfs(src, inf)
			if cur == 0 {
				break
			}
			flow += cur
		}
	}
	return flow
}

type Graph struct {
	node  []int
	next  []int
	to    []int
	flow  []int
	limit []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.flow = make([]int, e+3)
	g.limit = make([]int, e+3)
	g.cur = 1
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.limit[g.cur] = w
	g.flow[g.cur] = 0
}
