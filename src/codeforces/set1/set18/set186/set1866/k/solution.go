package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"math/big"
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
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	q := readNum(reader)
	qs := make([][]int, q)
	for i := range q {
		qs[i] = readNNums(reader, 2)
	}
	return solve(n, edges, qs)
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	var ans int

	nodes := make([]struct{ fi, se, fiW int }, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		p := &nodes[v]
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			d := dfs(w, v) + e.wt
			ans = max(ans, p.fi+d)
			if d > p.fi {
				p.se = p.fi
				p.fi = d
				p.fiW = w
			} else if d > p.se {
				p.se = d
			}
		}
		return p.fi
	}
	dfs(0, -1)

	hulls := make([][2][]vec, n)
	var reroot func(int, int, vec)
	reroot = func(v, fa int, up vec) {
		a := make([]vec, len(g[v]))
		for i, e := range g[v] {
			w := e.to
			if w == fa {
				a[i] = up
			} else {
				a[i] = vec{e.wt, nodes[w].fi}
			}
		}

		f := func(a, b vec) int { return cmp.Or(a.x-b.x, a.y-b.y) }
		slices.SortFunc(a, f)
		q := a[:0]
		b := []vec{}
		for _, v := range a {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) >= 0 {
				b = append(b, q[len(q)-1])
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		hulls[v][0] = q

		slices.SortFunc(b, f)
		q = b[:0]
		for _, v := range b {
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(v.sub(q[len(q)-1])) >= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		hulls[v][1] = q

		p := nodes[v]
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			down := p.fi
			if w == p.fiW {
				down = p.se
			}
			reroot(w, v, vec{e.wt, max(up.x+up.y, down)})
		}
	}
	reroot(0, -1, vec{})

	res := make([]int, len(queries))

	for i, cur := range queries {
		v := cur[0]
		p := vec{cur[1], 1}
		h := hulls[v-1][0]
		j := sort.Search(len(h)-1, func(j int) bool { return p.dot(h[j]) > p.dot(h[j+1]) })
		mx := p.dot(h[j])
		mx2 := 0
		if j > 0 {
			mx2 = p.dot(h[j-1])
		}
		if j < len(h)-1 {
			mx2 = max(mx2, p.dot(h[j+1]))
		}
		h = hulls[v-1][1]
		if len(h) > 0 {
			j := sort.Search(len(h)-1, func(j int) bool { return p.dot(h[j]) > p.dot(h[j+1]) })
			mx2 = max(mx2, p.dot(h[j]))
		}
		res[i] = max(mx+mx2, ans)
	}
	return res
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

func (g *Graph) AddEdge(u, v int, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
