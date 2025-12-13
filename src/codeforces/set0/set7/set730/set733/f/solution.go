package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, _, _, _, best, res := drive(reader)
	fmt.Fprintln(writer, best)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (n int, S int, W []int, C []int, edges [][]int, best int, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	W = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &W[i])
	}
	C = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &C[i])
	}
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	fmt.Fscan(reader, &S)
	best, res = solve(n, S, W, C, edges)
	return
}

type Edge struct {
	u  int
	v  int
	w  int
	c  int
	id int
}

func solve(n int, S int, W []int, C []int, edges [][]int) (int, [][]int) {
	m := len(edges)
	arr := make([]Edge, m)
	for i := range m {
		arr[i] = Edge{edges[i][0] - 1, edges[i][1] - 1, W[i], C[i], i}
	}

	slices.SortFunc(arr, func(a, b Edge) int {
		return cmp.Or(a.w-b.w, a.id-b.id)
	})

	set := NewDSU(n)

	g := NewGraph(n, 2*n)
	h := 20
	fa := make([][]int, n)
	dp := make([][]int, n)
	dep := make([]int, n)
	for i := range n {
		fa[i] = make([]int, h)
		dp[i] = make([]int, h)
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
			dp[u][i] = max(dp[u][i-1], dp[fa[u][i-1]][i-1])
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			eid := g.val[i]
			// w := W[eid]
			// c := C[eid]
			if p != v {
				dep[v] = dep[u] + 1
				dp[v][0] = W[eid]
				fa[v][0] = u

				dfs(u, v)
			}
		}
	}

	var sum int
	for i := range m {
		u, v := arr[i].u, arr[i].v
		if set.Union(u, v) {
			sum += arr[i].w
			g.AddEdge(u, v, arr[i].id)
			g.AddEdge(v, u, arr[i].id)
		}
	}

	dfs(0, 0)

	eid := -1
	best := 1 << 60

	find := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		var w int
		for i := h - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				w = max(w, dp[u][i])
				u = fa[u][i]
			}
		}
		if u == v {
			return w
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				w = max(w, dp[u][i], dp[v][i])
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return max(w, dp[u][0], dp[v][0])
	}

	for i := range m {
		u, v := arr[i].u, arr[i].v
		tmp := sum - find(u, v) + arr[i].w - S/arr[i].c
		if eid < 0 || tmp < best {
			best = tmp
			eid = i
		}
	}

	arr[eid].w -= S / arr[eid].c
	slices.SortFunc(arr, func(a, b Edge) int {
		return cmp.Or(a.w-b.w, a.id-b.id)
	})

	set.Reset()
	var res [][]int

	for i := range m {
		u, v := arr[i].u, arr[i].v
		if set.Union(u, v) {
			res = append(res, []int{arr[i].id + 1, arr[i].w})
		}
	}

	return best, res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := range n {
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

func (this *DSU) Reset() {
	for i := range this.arr {
		this.arr[i] = i
		this.cnt[i] = 1
	}
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
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
