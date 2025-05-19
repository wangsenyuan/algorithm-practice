package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	k := readNum(reader)
	queries := make([][]int, k)
	for i := range k {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, edges, queries)
}

const MOD = 1e9 + 7

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	id := make([]int, n)
	var m int
	comp := make([]int, n)
	dfn := make([]int, n)
	low := make([]int, n)
	vis := make([]bool, n)
	stack := make([]int, n)

	var top int
	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		vis[u] = true
		stack[top] = u
		top++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if !vis[v] {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else {
				low[u] = min(low[u], dfn[v])
			}
		}
		if dfn[u] == low[u] {
			for top > 0 {
				v := stack[top-1]
				top--
				id[v] = m
				comp[m]++
				vis[v] = false
				if v == u {
					break
				}
			}
			m++
		}
	}

	dfs(-1, 0)

	tr := NewGraph(m, 2*m)
	set := NewDSU(m)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if set.Union(id[u], id[v]) {
			tr.AddEdge(id[u], id[v])
			tr.AddEdge(id[v], id[u])
		}
	}

	sum := make([]int, m)
	dep := make([]int, m)
	fa := make([][]int, m)
	h := bits.Len(uint(m))

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		if comp[u] > 1 {
			sum[u]++
		}
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := tr.nodes[u]; i > 0; i = tr.next[i] {
			v := tr.to[i]
			if p != v {
				sum[v] = sum[u]
				dep[v] = dep[u] + 1
				dfs2(u, v)
			}
		}
	}
	dfs2(0, 0)

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

	pw := make([]int, m+1)
	pw[0] = 1
	for i := 1; i <= m; i++ {
		pw[i] = pw[i-1] * 2 % MOD
	}

	find := func(u int, v int) int {
		u = id[u]
		v = id[v]
		p := lca(u, v)
		tmp := sum[u] + sum[v] - 2*sum[p]
		if comp[p] > 1 {
			tmp++
		}
		return pw[tmp]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		ans[i] = find(u, v)
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
	next := make([]int, e+10)
	to := make([]int, e+10)
	return &Graph{nodes, next, to, 0}
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
