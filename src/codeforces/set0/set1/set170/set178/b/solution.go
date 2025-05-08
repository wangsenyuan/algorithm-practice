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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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
	merchants := make([][]int, k)
	for i := range k {
		merchants[i] = readNNums(reader, 2)
	}
	return solve(n, edges, merchants)
}

func solve(n int, roads [][]int, merchants [][]int) []int {
	m := len(roads)
	g := NewGraph(n, 2*m)
	for i, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)

	}

	dfn := make([]int, n)
	low := make([]int, n)
	marked := make([]bool, n)
	stack := make([]int, n)

	var top int
	var timer int

	val := make([]int, m)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		marked[u] = true
		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if p == v {
				continue
			}
			if dfn[v] == 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
				if low[v] > dfn[u] {
					// u -> v 是一个bridge
					val[j] = 1
				}
			} else if marked[v] {
				low[u] = min(low[u], dfn[v])
			}
		}
	}
	// 图是联通的
	dfs(-1, 0)

	g2 := NewGraph(n, 2*n)
	set := NewDSU(n)
	for i, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		if set.Union(u, v) {
			g2.AddEdge(u, v, val[i])
			g2.AddEdge(v, u, val[i])
		}
	}

	h := bits.Len(uint(n))
	fa := make([][]int, n)
	dep := make([]int, n)
	sum := make([]int, n)

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				sum[v] = sum[u] + g2.val[i]
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

	ans := make([]int, len(merchants))

	for i, cur := range merchants {
		s, l := cur[0]-1, cur[1]-1
		if s != l {
			ans[i] = sum[s] + sum[l] - 2*sum[lca(s, l)]
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
