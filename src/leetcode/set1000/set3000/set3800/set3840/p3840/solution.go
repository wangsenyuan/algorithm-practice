package p3840

import (
	"math/bits"
	"strconv"
	"strings"
)

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)
	in := make([]int, n)
	dep := make([]int, n)

	h := bits.Len(uint(n)) + 1
	fa := make([][]int, n)

	var timer int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		in[u] = timer
		timer++
		sz[u]++
		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(0, 0)

	tr := make(SegTree, 2*n)

	a := make([]int, n)

	for i := range n {
		x := int(s[i] - 'a')
		tr.Update(in[i], in[i]+sz[i], 1<<x)
		a[i] = x
	}

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

	var res []bool

	for _, cur := range queries {
		ss := strings.Split(cur, " ")
		if strings.HasPrefix(cur, "update") {
			u, _ := strconv.Atoi(ss[1])
			x := int(ss[2][0] - 'a')
			tr.Update(in[u], in[u]+sz[u], (1<<a[u])^(1<<x))
			a[u] = x
		} else {
			// query
			l, _ := strconv.Atoi(ss[1])
			r, _ := strconv.Atoi(ss[2])
			sum := tr.Get(in[l]) ^ tr.Get(in[r])
			// 还需要知道lcp
			p := lca(l, r)
			sum ^= 1 << a[p]
			d := bits.OnesCount(uint(sum))
			res = append(res, d <= 1)
		}
	}

	return res
}

type SegTree []int

func (t SegTree) Update(l int, r int, v int) {
	l += len(t) / 2
	r += len(t) / 2
	for l < r {
		if l&1 == 1 {
			t[l] ^= v
			l++
		}
		if r&1 == 1 {
			r--
			t[r] ^= v
		}
		l >>= 1
		r >>= 1
	}
}

func (t SegTree) Get(p int) int {
	p += len(t) / 2
	var res int
	for p > 0 {
		res ^= t[p]
		p >>= 1
	}
	return res
}
