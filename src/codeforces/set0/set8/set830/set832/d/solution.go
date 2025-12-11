package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, p, queries)
}

func solve(n int, p []int, queries [][]int) []int {
	g := NewGraph(n, n)

	for i := 1; i < n; i++ {
		g.AddEdge(p[i-1]-1, i)
	}

	h := bits.Len(uint(n)) + 1

	fa := make([][]int, n)

	dep := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		fa[u] = make([]int, h)
		if u > 0 {
			fa[u][0] = p[u-1] - 1
		}
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(v)
		}
	}

	dfs(0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
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

	dist := func(s int, t int, f int) int {
		u := lca(s, f)
		w := lca(u, t)
		if w == u {
			// 如果t在u子树的内部
			v1 := lca(t, f)
			v2 := lca(t, s)
			if v2 == w {
				return dep[f] - dep[v1] + 1
			}
			// v1 == w
			return dep[f] - dep[w] + dep[v2] - dep[w] + 1
		}
		return dep[f] - dep[u] + 1
	}

	find := func(a int, b int, c int) int {
		// 一共6种组合
		arr := []int{a, b, c}
		var res int
		for i := range 3 {
			for j := range 3 {
				if i != j {
					k := 3 - i - j
					d := dist(arr[i], arr[j], arr[k])
					res = max(res, d)
				}
			}
		}
		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0]-1, cur[1]-1, cur[2]-1)
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
