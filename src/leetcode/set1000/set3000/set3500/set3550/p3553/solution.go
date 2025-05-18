package p3553

import "math/bits"

func minimumWeight(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dep := make([]int, n)
	sum := make([]int, n)
	h := bits.Len(uint(n))
	fa := make([][]int, n)
	for i := range fa {
		fa[i] = make([]int, h)
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dep[v] = dep[u] + 1
				sum[v] = sum[u] + g.val[i]
				dfs(u, v)
			}
		}
	}
	dfs(0, 0)

	lca := func(u, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[fa[u][i]] >= dep[v] {
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

	find := func(a, b, c int) int {
		ab := lca(a, b)
		bc := lca(b, c)
		ac := lca(a, c)
		abc := lca(ab, c)
		if dep[ab] > dep[bc] {
			ans := sum[a] + sum[b] - 2*sum[ab]
			ans += sum[ab] + sum[c] - 2*sum[abc]
			return ans
		}
		if dep[ac] > dep[bc] {
			ans := sum[a] + sum[c] - 2*sum[ac]
			ans += sum[ac] + sum[b] - 2*sum[ab]
			return ans
		}
		// bc <= ac
		ans := sum[b] + sum[c] - 2*sum[bc]
		ans += sum[bc] + sum[a] - 2*sum[ab]
		return ans
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = find(q[0], q[1], q[2])
	}
	return ans
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
