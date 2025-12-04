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
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, r int
	fmt.Fscan(reader, &n, &r)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := range m {
		var p, q int
		fmt.Fscan(reader, &p, &q)
		queries[i] = []int{p, q}
	}
	return solve(r, a, edges, queries)
}

const inf = 1 << 60

func solve(root int, a []int, edges [][]int, queries [][]int) []int {
	n := len(a)

	root--

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var timer int
	in := make([]int, n)
	sz := make([]int, n)

	dep := make([]int, n)

	var dfs func(p int, u int, d int) int

	dfs = func(p int, u int, d int) int {
		dep[u] = d
		in[u] = timer
		timer++
		sz[u] = 1
		md := d
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				md = max(md, dfs(u, v, d+1))
				sz[u] += sz[v]
			}
		}
		return md
	}

	md := dfs(-1, root, 0)

	trs := make([]*node, md+1)
	trs[0] = build(0, n-1)

	que := make([]int, n)
	var head, tail int
	que[head] = root
	head++
	for tail < head {
		u := que[tail]
		tail++

		d := dep[u]
		if trs[d] != nil {
			trs[d] = trs[d].add(in[u], a[u])
		} else {
			trs[d] = trs[d-1].add(in[u], a[u])
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dep[v] == d+1 {
				que[head] = v
				head++
			}
		}
	}

	find := func(x int, k int) int {
		w := min(md, dep[x]+k)
		return trs[w].query(in[x], in[x]+sz[x]-1)
	}

	ans := make([]int, len(queries))

	var last int

	for i, cur := range queries {
		p, q := cur[0], cur[1]
		x := (p + last) % n
		k := (q + last) % n

		ans[i] = find(x, k)

		last = ans[i]
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

type node struct {
	lf, rg *node
	l, r   int
	val    int
}

func build(l int, r int) *node {
	node := &node{l: l, r: r, val: inf}

	if l < r {
		mid := (l + r) >> 1
		node.lf = build(l, mid)
		node.rg = build(mid+1, r)
	}

	return node
}

func (n *node) maintain() {
	n.val = min(n.lf.val, n.rg.val)
}

func (n node) add(p int, v int) *node {
	if n.l == n.r {
		n.val = v
	} else {
		mid := (n.l + n.r) >> 1
		if p <= mid {
			n.lf = (n.lf.add(p, v))
		} else {
			n.rg = (n.rg.add(p, v))
		}
		n.maintain()
	}

	return &n
}

func (n *node) query(L int, R int) int {
	if n.l == L && n.r == R {
		return n.val
	}
	res := inf
	mid := (n.l + n.r) >> 1
	if L <= mid {
		res = min(res, n.lf.query(L, min(mid, R)))
	}
	if mid < R {
		res = min(res, n.rg.query(max(mid+1, L), R))
	}
	return res
}
