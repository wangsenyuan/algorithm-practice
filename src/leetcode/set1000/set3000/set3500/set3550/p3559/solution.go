package p3559

import "math/bits"

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n)
	fa := make([][]int, n)
	h := bits.Len(uint(n))

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

	const mod = 1e9 + 7
	pw := make([]int, n)
	pw[0] = 1
	for i := 1; i < n; i++ {
		pw[i] = pw[i-1] * 2 % mod
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		if u == v {
			continue
		}
		p := lca(u, v)
		k := dep[u] + dep[v] - 2*dep[p]
		ans[i] = pw[k-1]
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
