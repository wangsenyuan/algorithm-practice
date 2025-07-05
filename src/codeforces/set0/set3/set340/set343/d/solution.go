package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	q := readNum(reader)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	big := make([]int, n)
	sz := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1
		big[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
			}
		}
	}
	dfs(-1, 0)

	head := make([]int, n)
	in := make([]int, n)
	pos := make([]int, n)
	var timer int

	fa := make([]int, n)

	var dfs2 func(p int, u int, x int)
	dfs2 = func(p int, u int, x int) {
		fa[u] = p
		head[u] = x
		in[u] = timer
		pos[timer] = u
		timer++
		if big[u] >= 0 {
			dfs2(u, big[u], x)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || big[u] == v {
				continue
			}
			dfs2(u, v, v)
		}
	}

	dfs2(-1, 0, 0)

	tr1 := make(SegTree, 2*n)
	tr2 := make(SegTree, 2*n)

	var ans []int

	findFillWaterAncestor := func(v int) int {
		var res int

		for v >= 0 {
			u := head[v]
			res = max(res, tr1.Get(in[u], in[v]+1))
			v = fa[u]
		}

		return res
	}

	findEmptyWaterDescendant := func(v int) int {
		i := in[v]
		j := in[v] + sz[v]
		return tr2.Get(i, j)
	}

	for i, cur := range queries {
		op, v := cur[0], cur[1]-1
		if op == 1 {
			tr1.Update(in[v], i+1)
		} else if op == 2 {
			tr2.Update(in[v], i+1)
		} else {
			a := findFillWaterAncestor(v)
			b := findEmptyWaterDescendant(v)
			if a > b {
				ans = append(ans, 1)
			} else {
				ans = append(ans, 0)
			}
		}
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

type SegTree []int

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v
	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
