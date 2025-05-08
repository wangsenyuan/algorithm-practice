package main

import (
	"bufio"
	"bytes"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	q := readNum(reader)
	qs := make([][]int, q)
	for i := range q {
		qs[i] = readNNums(reader, 2)
	}
	return solve(n, edges, qs)
}

func solve(n int, edges [][]int, queries [][]int) []bool {
	m := len(edges)

	g := NewGraph(n, 2*m)

	for i, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	comp := make([]int, n)
	for i := range n {
		comp[i] = -1
	}
	dep := make([]int, n)
	fa := make([][]int, n)
	h := bits.Len(uint(n))

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for j := 1; j < h; j++ {
			fa[u][j] = fa[fa[u][j-1]][j-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if comp[v] < 0 {
				dep[v] = dep[u] + 1
				comp[v] = comp[u]
				dfs(u, v)
			}
		}
	}

	stack := make([]int, m)
	top := 0

	check := func(k int) bool {
		tmp := top

		for tmp > 0 {
			i := stack[tmp-1]
			tmp--
			u, v := edges[i][0]-1, edges[i][1]-1
			if dep[u]%2 == dep[v]%2 {
				return true
			}
			if i == k {
				break
			}
		}
		top = tmp
		return false
	}

	val := make([]int, n)
	var timer int
	dfn := make([]int, n)
	low := make([]int, n)
	var tarjan func(p int, u int)
	tarjan = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if v == p || dfn[v] >= dfn[u] {
				// v may be u
				continue
			}

			stack[top] = w
			top++
			if dfn[v] == 0 {
				tarjan(u, v)
				low[u] = min(low[u], low[v])
				if low[v] >= dfn[u] && top > 0 && check(w) {
					for top > 0 {
						j := stack[top-1]
						top--
						x, y := edges[j][0]-1, edges[j][1]-1
						val[x] = 1
						val[y] = 1

						if j == w {
							val[u] = 0
							break
						}
					}
				}
			} else {
				// 这是一条回边
				low[u] = min(low[u], dfn[v])
			}
		}
	}

	marked := make([]bool, n)
	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		if p >= 0 {
			val[u] += val[p]
		}
		marked[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !marked[v] {
				dfs2(u, v)
			}
		}
	}

	for u := range n {
		if comp[u] < 0 {
			comp[u] = u
			dfs(u, u)
			tarjan(-1, u)
			dfs2(-1, u)
		}
	}

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

	get := func(u int, v int) bool {
		if comp[u] != comp[v] {
			return false
		}
		if dep[u]%2 != dep[v]%2 {
			return true
		}
		p := lca(u, v)
		return val[u]+val[v] > 2*val[p]
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		ans[i] = get(u, v)
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
