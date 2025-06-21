package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	m := readNum(reader)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries []string) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dfn := make([]int, n)
	ord := make([]int, 0, n)
	dep := make([]int, n)
	fa := make([][]int, n)
	h := bits.Len(uint(n))
	dist := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		dfn[u] = len(ord)
		ord = append(ord, u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dist[v] = dist[u] + g.val[i]
				dfs(u, v)
			}
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

	set := NewSet(n)

	var sum int

	get := func(u int, v int) int {
		tmp := dist[u] + dist[v] - 2*dist[lca(u, v)]
		return tmp
	}

	do := func(d int, m int) {
		x0 := set.LowerBound(n)
		if x0 < 0 {
			// 空集
			return
		}
		prev := set.LowerBound(d)
		if prev < 0 {
			prev = x0
		}
		nxt := set.UpperBound(d)
		if nxt == n {
			nxt = set.UpperBound(0)
		}

		sum += m * (get(ord[prev], ord[d]) + get(ord[d], ord[nxt]) - get(ord[prev], ord[nxt]))
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == '?' {
			ans = append(ans, sum/2)
			continue
		}
		var x int
		readInt([]byte(cur), 2, &x)
		x--
		d := dfn[x]
		if cur[0] == '+' {
			do(d, 1)
			set.Set(d)
		} else {
			set.Unset(d)
			do(d, -1)
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

type Set struct {
	min_pos *SegTree
	max_pos *SegTree
	n       int
}

func NewSet(n int) Set {
	min_pos := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})
	max_pos := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})
	return Set{min_pos, max_pos, n}
}

func (set *Set) Set(p int) {
	set.min_pos.Update(p, p)
	set.max_pos.Update(p, p)
}

func (set *Set) Unset(p int) {
	set.min_pos.Update(p, set.n)
	set.max_pos.Update(p, -1)
}

func (set Set) LowerBound(p int) int {
	return set.max_pos.Get(0, p)
}

func (set Set) UpperBound(p int) int {
	return set.min_pos.Get(p, set.n)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
