package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	ans := drive(reader)
	for _, x := range ans {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

type Edge struct {
	u  int
	v  int
	w  int
	id int
}

func solve(n int, edges [][]int) []int {
	m := len(edges)
	arr := make([]Edge, m)
	for i := range m {
		arr[i] = Edge{edges[i][0] - 1, edges[i][1] - 1, edges[i][2], i}
	}

	slices.SortFunc(arr, func(a, b Edge) int {
		return cmp.Or(a.w-b.w, a.id-b.id)
	})

	set := NewDSU(n)

	g := NewGraph(n, 2*n)

	var sum int

	marked := make([]bool, m)

	for _, e := range arr {
		u, v := e.u, e.v
		if set.Union(u, v) {
			g.AddEdge(u, v, e.w)
			g.AddEdge(v, u, e.w)
			marked[e.id] = true
			sum += e.w
		}
	}

	h := bits.Len(uint(n))
	fa := make([][]int, n)
	dp := make([][]int, n)

	dep := make([]int, n)

	var dfs func(p int, u int, w int)
	dfs = func(p int, u int, w int) {
		fa[u] = make([]int, h)
		dp[u] = make([]int, h)
		fa[u][0] = p
		dp[u][0] = w
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
			dp[u][i] = max(dp[u][i-1], dp[fa[u][i-1]][i-1])
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v, g.val[i])
			}
		}
	}

	dfs(0, 0, 0)

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

	findMaxEdge := func(u int, d int) int {
		var res int
		for i := h - 1; i >= 0; i-- {
			if d&(1<<i) > 0 {
				res = max(res, dp[u][i])
				u = fa[u][i]
			}
		}
		return res
	}

	ans := make([]int, m)
	for i, cur := range edges {
		if marked[i] {
			ans[i] = sum
		} else {
			u, v, w := cur[0]-1, cur[1]-1, cur[2]
			p := lca(u, v)
			var w1 int
			if u != p {
				w1 = max(w1, findMaxEdge(u, dep[u]-dep[p]))
			}
			if v != p {
				w1 = max(w1, findMaxEdge(v, dep[v]-dep[p]))
			}
			ans[i] = sum - w1 + w
		}
	}
	return ans
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
