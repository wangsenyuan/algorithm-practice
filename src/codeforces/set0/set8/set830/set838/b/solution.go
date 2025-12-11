package main

import (
	"bufio"
	"fmt"
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
	edges := make([][]int, 2*n-2)
	for i := range edges {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)

	back := make([]int, n)

	for i := range n - 1 {
		e := edges[i]
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	for i := n - 1; i < len(edges); i++ {
		back[edges[i][0]-1] = i
	}

	pos := make([]int, n)
	var timer int
	sz := make([]int, n)
	var dfs func(p int, u int)

	val := make([]int, n)

	dfs = func(p int, u int) {
		sz[u]++
		pos[u] = timer
		timer++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				val[v] = g.val[i]
				sz[u] += sz[v]
			}
		}
	}

	dfs(0, 0)

	tr := make(SegTree, 2*n)
	tr2 := NewTree(n)

	for u := range n {
		tr.Update(pos[u], pos[u]+sz[u], val[u])
		tr2.Update(pos[u], pos[u]+sz[u]-1, val[u])
	}

	for i := range n {
		if i > 0 {
			tr2.Update(pos[i], pos[i], edges[back[i]][2])
		}
	}

	isAnc := func(u int, v int) bool {
		return pos[u] < pos[v] && pos[v] < pos[u]+sz[u]
	}

	change := func(i int, w int) {
		if i >= n-1 {
			// a back edge
			// edges[i][2] = w
			delta := w - edges[i][2]
			v := edges[i][0] - 1

			tr2.Update(pos[v], pos[v], delta)

			edges[i][2] = w

			return
		}
		// i < n
		u, v := edges[i][0]-1, edges[i][1]-1
		if isAnc(v, u) {
			v = u
		}
		tr.Update(pos[v], pos[v]+sz[v], w-val[v])

		tr2.Update(pos[v], pos[v]+sz[v]-1, w-val[v])

		val[v] = w
	}

	find := func(u int, v int) int {
		if u == v {
			return 0
		}
		if isAnc(u, v) {
			return tr.Get(pos[v]) - tr.Get(pos[u])
		}
		// 需要知道back的edge
		// 这个地方不一定，需要知道，整个子树中，去root的最短距离
		res := tr.Get(pos[v])
		// 在子树u中，找到最小的到root的值
		w := tr2.Get(pos[u], pos[u]+sz[u]-1)
		w -= tr.Get(pos[u])
		res += w

		return res
	}

	var res []int

	for _, cur := range queries {
		if cur[0] == 1 {
			change(cur[1]-1, cur[2])
		} else {
			u, v := cur[1]-1, cur[2]-1
			res = append(res, find(u, v))
		}
	}

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

type SegTree []int

func (tr SegTree) Update(l int, r int, v int) {
	// range update, point get
	n := len(tr) / 2
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			tr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr SegTree) Get(p int) int {
	var res int
	n := len(tr) / 2
	p += n
	for p > 0 {
		res += tr[p]
		p >>= 1
	}
	return res
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) apply(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(i*2+1, t.lazy[i])
		t.apply(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = min(t.val[i*2+1], t.val[i*2+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) / 2
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		t.pull(i)
	}
	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return t.val[i]
		}
		t.push(i)

		mid := (l + r) / 2

		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}

		return min(f(i*2+1, l, mid, L, mid), f(i*2+2, mid+1, r, mid+1, R))
	}
	return f(0, 0, t.sz-1, L, R)
}
