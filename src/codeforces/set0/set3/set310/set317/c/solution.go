package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, ok, res := process(reader)
	if !ok {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) (V int, a []int, b []int, edges [][]int, ok bool, res [][]int) {
	var n, e int
	fmt.Fscan(reader, &n, &V, &e)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	edges = make([][]int, e)
	for i := range e {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	ok, res = solve(V, slices.Clone(a), b, edges)
	return
}

func solve(V int, a []int, b []int, edges [][]int) (bool, [][]int) {
	n := len(a)
	set := CreateUFSet(n)
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if set.Union(u, v) {
			g.AddEdge(u, v)
			g.AddEdge(v, u)
		}
	}

	type pair struct {
		first  int
		second int
	}

	var todo []pair
	fa := make([]int, n)

	bal := make([]int, n)
	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		fa[u] = p
		bal[u] = a[u] - b[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				bal[u] += dfs(u, v)
			}
		}
		if a[u] < b[u] {
			todo = append(todo, pair{b[u] - a[u], u})
		}
		return bal[u]
	}

	var res [][]int

	var pull func(s int, u int, x int)

	// s肯定有x的流量可以被使用
	pull = func(s int, u int, x int) {
		if u == s {
			return
		}
		p := fa[u]
		d := min(x, a[p])

		if d > 0 {
			a[p] -= d
			a[u] += d
			res = append(res, []int{p + 1, u + 1, d})
		}
		pull(s, p, x)
		if d < x {
			a[p] -= x - d
			a[u] += x - d
			res = append(res, []int{p + 1, u + 1, x - d})
		}
	}

	for u := 0; u < n; u++ {
		if a[u] > b[u] {
			todo = todo[:0]
			tot := dfs(-1, u)
			if tot != 0 {
				return false, nil
			}

			d := a[u] - b[u]

			for i := 0; i < len(todo) && d > 0; i++ {
				x := min(d, todo[i].first)
				pull(u, todo[i].second, x)
				d -= x
			}

		}
	}

	for u := 0; u < n; u++ {
		if a[u] != b[u] {
			return false, nil
		}
	}

	return true, res
}

type UFSet struct {
	set  []int
	cnt  []int
	size int
}

func CreateUFSet(n int) UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
		set[i] = i
	}
	return UFSet{set, cnt, n}
}

func (uf *UFSet) Find(a int) int {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}
	return find(a)
}

func (uf *UFSet) Union(a, b int) bool {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}

	pa := find(a)
	pb := find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		uf.set[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	} else {
		uf.set[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	}
	uf.size--
	return true
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
