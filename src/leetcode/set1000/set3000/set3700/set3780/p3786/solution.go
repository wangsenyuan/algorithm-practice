package p3786

import (
	"math/bits"
	"slices"
)

func interactionCosts1(n int, edges [][]int, group []int) int64 {
	adj := make([][]int, n)
	x := slices.Max(group)
	cnt := make([]int, x+1)
	for _, y := range group {
		cnt[y]++
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var res int
	var dfs func(p int, u int)

	sz := make([][]int, x+1)
	for i := range x + 1 {
		sz[i] = make([]int, n)
	}

	dfs = func(p int, u int) {
		sz[group[u]][u]++
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)

				for k := range x + 1 {
					res += (cnt[k] - sz[k][v]) * sz[k][v]
					sz[k][u] += sz[k][v]
				}
			}
		}
	}

	dfs(-1, 0)

	return int64(res)
}

func interactionCosts(n int, edges [][]int, group []int) int64 {
	mx := slices.Max(group)
	h := bits.Len(uint(n)) + 1
	dep := make([]int, n)
	fa := make([][]int, n)
	dfn := make([]int, n)
	nodes := make([][]int, mx+1)

	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dfn[u] = timer
		timer++
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		nodes[group[u]] = append(nodes[group[u]], u)

		for _, v := range adj[u] {
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

	vt := make([][]int, n)

	connect := func(u int, v int) {
		vt[u] = append(vt[u], v)
	}

	stack := make([]int, n)
	tin := make([]int, n)

	play := func(x int) int {
		ws := nodes[x]
		// slices.SortFunc(ws, func(u int, v int) int {
		// 	return dfn[u] - dfn[v]
		// })

		// st[0] = 0
		top := 1
		vt[0] = vt[0][:0]
		for _, v := range ws {
			tin[v] = x
			if v == 0 {
				continue
			}
			vt[v] = vt[v][:0]
			p := lca(v, stack[top-1])
			for top > 1 && dfn[p] <= dfn[stack[top-2]] {
				connect(stack[top-2], stack[top-1])
				top--
			}
			if p != stack[top-1] {
				vt[p] = vt[p][:0]
				connect(p, stack[top-1])
				stack[top-1] = p
			}
			stack[top] = v
			top++
		}

		for i := 1; i < top; i++ {
			connect(stack[i-1], stack[i])
		}

		var ans int

		var dfs2 func(u int) (sz int)
		dfs2 = func(u int) (sz int) {
			if tin[u] == x {
				sz++
			}
			for _, v := range vt[u] {
				tmp := dfs2(v)
				d := dep[v] - dep[u]
				ans += d * (len(nodes[x]) - tmp) * tmp
				sz += tmp
			}
			return
		}
		if tin[0] != x && len(vt[0]) == 1 {
			dfs2(vt[0][0])
		} else {
			dfs2(0)
		}
		return ans
	}

	var res int
	for i := range mx + 1 {
		res += play(i)
	}

	return int64(res)
}
