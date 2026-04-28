package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	parent := make([]int, n-1)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &parent[i-1])
	}

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}

	return solve(parent, edges)
}

func solve(P []int, edges [][]int) int {
	n := len(P) + 1
	adj := make([][]int, n)
	parent := make([]int, n)
	for i := 1; i < n; i++ {
		parent[i] = P[i-1] - 1
		adj[parent[i]] = append(adj[parent[i]], i)
	}

	sz := make([]int, n)
	big := make([]int, n)
	dep := make([]int, n)
	var dfs1 func(u int)
	dfs1 = func(u int) {
		sz[u] = 1
		big[u] = -1
		for _, v := range adj[u] {
			dep[v] = dep[u] + 1
			dfs1(v)
			sz[u] += sz[v]
			if big[u] == -1 || sz[v] > sz[big[u]] {
				big[u] = v
			}
		}
	}

	dfs1(0)

	dfn := make([]int, n)
	var ord int

	rev := make([]int, n)

	head := make([]int, n)

	var dfs2 func(u int, x int)

	dfs2 = func(u int, x int) {
		head[u] = x
		dfn[u] = ord
		rev[ord] = u
		ord++
		if big[u] != -1 {
			dfs2(big[u], x)
		}
		for _, v := range adj[u] {
			if v != big[u] {
				dfs2(v, v)
			}
		}
	}

	dfs2(0, 0)

	lca := func(u int, v int) int {
		for head[u] != head[v] {
			// 移动head更深的那个
			if dep[head[u]] < dep[head[v]] {
				u, v = v, u
			}
			u = parent[head[u]]
		}
		if dep[u] > dep[v] {
			u = v
		}
		return u
	}

	todo := make([][]int, n)
	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		p := lca(u, v)
		todo[p] = append(todo[p], i)
	}

	// fp[u] 表示子树u的结果
	fp := make([]int, n)
	s := make([]int, n)

	tr := make(SegTree, 2*n)

	nextAncestor := func(u int, p int) int {
		for head[u] != head[p] {
			u = head[u]
			if parent[u] == p {
				return u
			}
			u = parent[u]
		}
		return rev[dfn[p]+1]
	}

	query := func(u int, p int) int {
		// p 肯定是u得祖先节点
		var res int
		for head[u] != head[p] {
			res += tr.Get(dfn[head[u]], dfn[u]+1)
			u = parent[head[u]]
		}
		// head[u] == head[p]
		res += tr.Get(dfn[p]+1, dfn[u]+1)
		return res
	}

	var dfs3 func(u int)
	dfs3 = func(u int) {
		for _, v := range adj[u] {
			dfs3(v)
			s[u] += fp[v]
		}

		// 要计算出fp[u]以后，才能计算
		best := s[u]
		for _, i := range todo[u] {
			x, y, c := edges[i][0]-1, edges[i][1]-1, edges[i][2]
			tmp := c + s[u]
			if x != u {
				tmp += s[x]
				x1 := nextAncestor(x, u)
				tmp += query(x, x1) - fp[x1]
			}
			if y != u {
				tmp += s[y]
				y1 := nextAncestor(y, u)
				tmp += query(y, y1) - fp[y1]
			}
			best = max(best, tmp)
		}
		fp[u] = best

		for _, v := range adj[u] {
			tr.Update(dfn[v], s[u]-fp[v])
		}
	}

	dfs3(0)

	return fp[0]
}

type SegTree []int

func (t SegTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	t[p] = v
	for p > 1 {
		t[p>>1] = t[p] + t[p^1]
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res += t[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
