package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	dep := make([]int, n)
	fa := make([][]int, n)
	h := bits.Len(uint(n))
	sz := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}
	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	kth := func(u int, k int) int {
		for i := h - 1; i >= 0; i-- {
			if k&(1<<i) > 0 {
				u = fa[u][i]
			}
		}
		return u
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		if u == v {
			ans[i] = n
			continue
		}
		p := lca(u, v)
		d := dep[u] + dep[v] - 2*dep[p]
		if d%2 == 1 {
			// 无法找到一个距离相等的房间
			ans[i] = 0
			continue
		}
		d /= 2

		if dep[u] == dep[v] {
			// p is the mid
			ans[i] = n - sz[kth(u, d-1)] - sz[kth(v, d-1)]
			continue
		}

		if dep[u] < dep[v] {
			u = v
		}
		// u是那个离p更远的点
		mid := kth(u, d)
		ans[i] = sz[mid] - sz[kth(u, d-1)]
	}

	return ans
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
