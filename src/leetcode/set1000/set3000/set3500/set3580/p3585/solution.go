package p3585

import "math/bits"

func findMedian(n int, edges [][]int, queries [][]int) []int {
	// 如果p = lca(u, v)
	// 如果p是x返回p
	// 如果p不是，两种情况，一种是x在u,p中间，一种是,x在v,p中间
	// 那这个时候可以使用二分去查
	g := NewGraph(n, len(edges)*2)
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
		g.AddEdge(e[1], e[0], e[2])
	}

	dep := make([]int, n)
	sum := make([]int, n)
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
				sum[v] = sum[u] + g.val[i]
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

	find2 := func(u int, val int) int {
		i := h - 1
		for i >= 0 && (sum[u]-sum[fa[u][i]])*2 > val {
			i--
		}

		if i < 0 {
			return u
		}

		v := fa[u][i]
		// i >= 0
		// (sum[u] - sum[fa[u][i]]) * 2 < val
		for i > 0 {
			i--
			if (sum[u]-sum[fa[v][i]])*2 <= val {
				v = fa[v][i]
			}
		}
		return v
	}

	find := func(u int, v int) int {
		if u == v {
			return u
		}
		p := lca(u, v)
		tmp1 := sum[u] - sum[p] + sum[v] - sum[p]
		if (sum[u]-sum[p])*2 >= tmp1 {
			x := find2(u, tmp1)
			if (sum[u]-sum[x])*2 < tmp1 {
				x = fa[x][0]
			}
			return x
		}
		// x出现在路径(p -> v)
		// dist(x, v) * 2 <= tmp1
		return find2(v, tmp1)
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = find(q[0], q[1])
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
