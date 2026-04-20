package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)
	big := make([]int, n)
	height := make([]int, n)
	fa := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		big[u] = -1
		sz[u] = 1
		fa[u] = p
		for _, v := range adj[u] {
			if p != v {
				height[v] = height[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
			}
		}
	}

	dfs(-1, 0)

	head := make([]int, n)
	pos := make([]int, n)
	var timer int

	var dfs2 func(p int, u int, x int)
	dfs2 = func(p int, u int, x int) {
		head[u] = x
		pos[u] = timer
		timer++
		if big[u] >= 0 {
			dfs2(u, big[u], x)
		}
		for _, v := range adj[u] {
			if p != v && big[u] != v {
				dfs2(u, v, v)
			}
		}
	}

	dfs2(-1, 0, 0)

	buf := make([]int, n)
	for i := range n {
		buf[pos[i]] = height[i]
	}

	tr := NewTree(n, buf)

	update := func(u int) {
		h := height[u]
		for u >= 0 {
			w := head[u]
			tr.Update(pos[w], pos[u], h)
			u = fa[w]
		}
	}

	update(0)

	get := func(u int) int {
		res := inf
		// res = height[?] - 2 * height[lca] + height[u]
		h := height[u]
		for u >= 0 {
			w := head[u]
			res = min(res, tr.Get(pos[w], pos[u]))
			u = fa[w]
		}

		res += h

		return res
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			update(cur[1] - 1)
		} else {
			ans = append(ans, get(cur[1]-1))
		}
	}

	return ans
}

const inf = 1 << 60

type Tree struct {
	val  []int
	arr  []int
	lazy []int
	sz   int
}

func NewTree(n int, input []int) *Tree {
	val := make([]int, 4*n)
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)

	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		lazy[i] = inf
		if l == r {
			arr[i] = input[l]
			val[i] = inf
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		arr[i] = max(arr[i*2+1], arr[i*2+2])
	}

	f(0, 0, n-1)

	return &Tree{val, arr, lazy, n}
}

func (tr *Tree) apply(i int, v int) {
	if tr.lazy[i] > v {
		tr.lazy[i] = v
		tr.val[i] = min(tr.val[i], v-2*tr.arr[i])
	}
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != inf {
		tr.apply(i*2+1, tr.lazy[i])
		tr.apply(i*2+2, tr.lazy[i])
		tr.lazy[i] = inf
	}
}

func (tr *Tree) pull(i int) {
	tr.val[i] = min(tr.val[i*2+1], tr.val[i*2+2])
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) / 2
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.pull(i)
	}
	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int

	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i)
		mid := (l + r) / 2
		res := inf
		if L <= mid {
			res = f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res = min(res, f(i*2+2, mid+1, r, max(mid+1, L), R))
		}
		return res
	}

	return f(0, 0, tr.sz-1, L, R)
}
