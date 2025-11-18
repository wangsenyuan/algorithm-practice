package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	for _, row := range res {
		for _, v := range row {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (n int, a []int, b []int, edges [][]int, ok bool, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	edges = make([][]int, m)
	for i := range m {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	ok, res = solve(n, a, b, edges)
	return
}

func solve(n int, a []int, b []int, edges [][]int) (bool, [][]int) {
	m := len(edges)
	g := NewGraph(2*n+2, m*4+n*6+10)

	var src int
	snk := 2*n + 1

	addEdge := func(u, v, w int) {
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, 0)
	}

	var s1, s2 int

	for i := range n {
		addEdge(src, 2*i+1, a[i])
		addEdge(2*i+2, snk, b[i])
		addEdge(2*i+1, 2*i+2, inf)
		s1 += a[i]
		s2 += b[i]
	}

	if s1 != s2 {
		return false, nil
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		addEdge(2*u+1, 2*v+2, inf)
		addEdge(2*v+1, 2*u+2, inf)
	}

	sum := dinic(src, snk, 2*n+2, g)

	if s1 != sum {
		return false, nil
	}
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, n)
	}

	for i := range n {
		u := 2*i + 1
		for j := g.node[u]; j > 0; j = g.next[j] {
			v := g.to[j]
			if v != src && v != snk && v&1 == 0 {
				v = (v - 2) / 2
				if g.flow[j] > 0 {
					res[i][v] = g.flow[j]
				}
			}
		}
	}

	return true, res
}

const inf = 1 << 60

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
