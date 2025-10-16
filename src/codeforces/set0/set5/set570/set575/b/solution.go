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
	fmt.Println(res)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	roads := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		roads[i] = make([]int, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	var m int
	fmt.Fscan(reader, &m)
	stops := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &stops[i])
	}
	return solve(n, roads, stops)
}

func solve(n int, roads [][]int, stops []int) int {
	g := NewGraph(n, 2*n)

	for _, cur := range roads {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, -w)
	}

	dep := make([]int, n)
	h := bits.Len(uint(n))
	fa := make([][]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
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

	// 应该计算每条边的贡献, (a, b, 1) 那么就需要计算有多少次从b->a的通过
	// 假设a是b的父节点，在b中共出现了x次多余的，那么肯定是从a -> b 经过了x次
	// 但咋算呢？
	ends_down := make([]int, n)
	ends_up := make([]int, n)
	gone_up := make([]int, n)
	u := 0
	for _, v := range stops {
		v--
		gone_up[u]++
		ends_down[v]++
		p := lca(u, v)
		ends_up[p]++
		u = v
	}
	m := len(stops) + 1

	pw := make([]int, m+1)
	pw[0] = 1
	for i := 1; i <= m; i++ {
		pw[i] = pw[i-1] * 2 % mod
	}

	var res int

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
				ends_down[u] += ends_down[v]
				ends_up[u] += ends_up[v]
				gone_up[u] += gone_up[v]
				switch g.val[i] {
				case -1:
					// 那么从u到v需要付钱
					x := ends_down[v] - ends_up[v]
					res = add(res, sub(pw[x], 1))
				case 1:
					// 从v到u需要付钱
					x := gone_up[v] - ends_up[v]
					res = add(res, sub(pw[x], 1))
				}
			}
		}
	}

	dfs2(-1, 0)

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
