package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"math/bits"
	"os"
	"slices"
	"sort"
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
	a := readNNums(reader, n)
	f := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, f, edges, queries)
}

func solve(a []int, f []int, edges [][]int, queries [][]int) []int {
	f1 := sortAndUnique(f)
	n := len(f)
	for i := range n {
		f[i] = sort.SearchInts(f1, f[i])
	}

	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	in := make([]int32, n)
	out := make([]int32, n)
	var ord []int32

	var dfs func(p int, u int)

	h := bits.Len(uint(n))
	fa := make([][]int32, n)
	dep := make([]int, n)
	dfs = func(p int, u int) {
		fa[u] = make([]int32, n)
		fa[u][0] = int32(p)
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		in[u] = int32(len(ord))
		ord = append(ord, int32(u))
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
		out[u] = int32(len(ord))
		ord = append(ord, int32(u))
	}

	dfs(0, 0)

	lca := func(u int32, v int32) int32 {
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

	m := len(f1)

	type query struct {
		id int
		a  int32
		b  int32
		l  int32
		r  int32
	}

	isAnc := func(a, b int) bool {
		return in[a] < in[b] && out[b] < out[a]
	}

	qs := make([]query, len(queries))
	for i, q := range queries {
		a, b := q[0]-1, q[1]-1
		var l, r int32
		if isAnc(a, b) {
			l, r = in[a], in[b]
		} else if isAnc(b, a) {
			l, r = out[a], out[b]
		} else if out[a] < in[b] {
			l, r = out[a], in[b]
		} else {
			l, r = out[b], in[a]
		}
		qs[i] = query{i, int32(a), int32(b), l, r}
	}

	block_size := int32(math.Sqrt(float64(n*2))) + 1

	slices.SortFunc(qs, func(u query, v query) int {
		if u.r/block_size != v.r/block_size {
			return int(u.r - v.r)
		}
		d := u.r / block_size
		if d%2 == 0 {
			return int(u.l - v.l)
		} else {
			return int(v.l - u.l)
		}
	})

	ans := make([]int, len(queries))

	var sum int
	marked := make([]bool, n)
	boys := make([]int32, m)
	girls := make([]int32, m)

	doIt := func(u int32) {
		if !marked[u] {
			if a[u] == 0 {
				sum += int(boys[f[u]])
				girls[f[u]]++
			} else {
				sum += int(girls[f[u]])
				boys[f[u]]++
			}
		} else {
			if a[u] == 0 {
				sum -= int(boys[f[u]])
				girls[f[u]]--
			} else {
				sum -= int(girls[f[u]])
				boys[f[u]]--
			}
		}
		marked[u] = !marked[u]
	}

	var L, R int32 = 0, 0

	for _, q := range qs {
		for R <= q.r {
			doIt(ord[R])
			R++
		}

		for L > q.l {
			L--
			doIt(ord[L])
		}

		for R-1 > q.r {
			R--
			doIt(ord[R])
		}
		for L < q.l {
			doIt(ord[L])
			L++
		}

		ans[q.id] = sum

		c := int32(lca(int32(q.a), int32(q.b)))

		if c != q.a && c != q.b {
			if a[c] == 0 {
				ans[q.id] += int(boys[f[c]])
			} else {
				ans[q.id] += int(girls[f[c]])
			}
		}
	}

	return ans
}

func sortAndUnique(a []int) []int {
	arr := slices.Clone(a)
	sort.Ints(arr)
	return slices.Compact(arr)
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
