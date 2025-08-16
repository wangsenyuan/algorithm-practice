package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	res := solve(n, edges)
	s := fmt.Sprintf("%v", res)
	return s[1 : len(s)-1]
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([]int, n)
	for i := range n {
		dist[i] = -1
	}

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	dist[0] = 0
	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				que[head] = v
				head++
			}
		}
	}
	md := slices.Max(dist)
	var root int
	for u := range n {
		if dist[u] == md {
			root = u
			break
		}
	}

	pos := make([]int, n)
	sz := make([]int, n)
	fa := make([]int, n)
	dep := make([]int, n)
	var order []int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		sz[u] = 1
		pos[u] = len(order)
		order = append(order, u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, root)

	arr := make([]int, n)
	for u := range n {
		arr[u] = dep[order[u]]
	}

	tr := NewTree(arr)

	ans := make([]int, n)
	ans[0] = 1

	marked := make([]bool, n)
	marked[root] = true

	for i := 1; i < n; i++ {
		if ans[i-1] == n {
			ans[i] = n
			continue
		}
		far := tr.FindMaxPos()
		ans[i] = ans[i-1] + far.first
		j := order[far.second]

		for !marked[j] {
			marked[j] = true
			tr.Update(pos[j], pos[j]+sz[j]-1, -1)
			j = fa[j]
		}
	}

	return ans
}

type Tree struct {
	arr  []int
	lazy []int
}

func NewTree(arr []int) *Tree {
	n := len(arr)
	rra := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			rra[i] = arr[l]
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		rra[i] = max(rra[i*2+1], rra[i*2+2])
	}
	build(0, 0, n-1)
	return &Tree{rra, lazy}
}

func (tr *Tree) update(i int, v int) {
	tr.lazy[i] += v
	tr.arr[i] += v
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(i*2+1, tr.lazy[i])
		tr.update(i*2+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.update(i, v)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(L, mid+1), R)
		}
		tr.arr[i] = max(tr.arr[i*2+1], tr.arr[i*2+2])
	}
	f(0, 0, len(tr.arr)/4-1, L, R)
}

func (tr *Tree) FindMaxPos() pair {
	var f func(i int, l int, r int) pair
	f = func(i int, l int, r int) pair {
		if l == r {
			return pair{tr.arr[i], l}
		}
		tr.push(i)
		mid := (l + r) >> 1
		if tr.arr[i*2+1] >= tr.arr[i*2+2] {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, len(tr.arr)/4-1)
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
