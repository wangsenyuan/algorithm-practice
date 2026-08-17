package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		cost, colors := drive(reader)
		fmt.Fprintln(writer, cost)
		s := fmt.Sprintf("%v", colors)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (int64, []int) {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	w := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &w[i])
	}
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, k, w, c, edges)
}

func solve(n, k int, w, c []int, edges [][]int) (int64, []int) {

	if slices.Max(c) == 0 {
		assign := make([]int, n)
		for i := range n {
			assign[i] = 1
		}
		return 0, assign
	}

	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	h := bits.Len(uint(n))
	fa := make([][]int, n)
	dep := make([]int, n)
	dfn := make([]int, n)
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

	cols := make([][]int, k+1)
	for i := 0; i < n; i++ {
		if c[i] > 0 {
			cols[c[i]] = append(cols[c[i]], i)
		}
	}

	mark := make([]bool, n)

	cutie := make([]bool, n)

	cnt := make([]int, n)

	col := make([]int, n)

	for x := 1; x <= k; x++ {
		slices.SortFunc(cols[x], func(u int, v int) int {
			return dfn[u] - dfn[v]
		})

		var V []int

		for j := 0; j+1 < len(cols[x]); j++ {
			u, v := cols[x][j], cols[x][j+1]
			p := lca(u, v)
			if !mark[p] {
				if c[p] > 0 {
					if c[p] != x {
						cutie[p] = true
					}
					continue
				}
				mark[p] = true
				cnt[p]++
				col[p] = x
				V = append(V, p)
			}
		}
		for _, v := range V {
			mark[v] = false
		}
	}

	var sum int

	for u := range n {
		if cutie[u] || cnt[u] > 1 {
			sum += w[u]
		}
		if c[u] == 0 && cnt[u] == 1 {
			c[u] = col[u]
		}
	}

	var dfs2 func(u int, cc int)
	dfs2 = func(u int, cc int) {
		if c[u] > 0 {
			cc = c[u]
		} else {
			c[u] = cc
		}
		for _, v := range adj[u] {
			if v != fa[u][0] {
				dfs2(v, cc)
			}
		}
	}

	for u := range n {
		if c[u] > 0 {
			dfs2(0, c[u])
			break
		}
	}

	return int64(sum), c
}
