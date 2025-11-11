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
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintf(writer, "%d ", x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n-1)
	for i := range n - 1 {
		fmt.Fscan(reader, &p[i])
	}
	return solve(n, p)
}

func solve(n int, p []int) []int {
	g := NewGraph(n, 2*n)
	for i, x := range p {
		g.AddEdge(x-1, i+1)
	}

	h := bits.Len(uint(n))
	fa := make([][]int, n)
	dep := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(u, v)
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

	dist := func(u int, v int) int {
		p := lca(u, v)
		return dep[u] + dep[v] - 2*dep[p]
	}

	dia := make([][]int, n)

	// dia[i] = {u, v} 是i compnent的直径的两个端点
	for i := range n {
		dia[i] = []int{i, i, 0}
	}

	ans := make([]int, n-1)
	set := NewDSU(n)

	for i, x := range p {
		u := x - 1
		v := i + 1
		fu := set.Find(u)
		fv := set.Find(v)

		tmp := slices.Clone(dia[fu])

		if dia[fv][2] > tmp[2] {
			tmp = slices.Clone(dia[fv])
		}

		for l := range 2 {
			for r := range 2 {
				d := dist(dia[fu][l], dia[fv][r])
				if d > tmp[2] {
					tmp[0] = dia[fu][l]
					tmp[1] = dia[fv][r]
					tmp[2] = d
				}
			}
		}

		set.Union(fu, fv)
		fu = set.Find(fu)
		dia[fu] = tmp
		ans[i] = tmp[2]
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

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
