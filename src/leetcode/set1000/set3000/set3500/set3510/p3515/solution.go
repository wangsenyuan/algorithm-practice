package p3515

func treeQueries(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	in := make([]int, n)
	sz := make([]int, n)
	val := make([]int, n)

	tr := NewSegTree(n)
	fa := make([]int, n)

	var pos int
	var dfs func(p int, u int, sum int)
	dfs = func(p int, u int, sum int) {
		fa[u] = p
		in[u] = pos
		tr.Update(pos, pos+1, sum)
		pos++
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if p != v {
				val[v] = w
				dfs(u, v, sum+w)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0, 0)

	var res []int

	for _, cur := range queries {
		if cur[0] == 1 {
			u, v, w := cur[1], cur[2], cur[3]
			u--
			v--
			if fa[u] == v {
				v = u
			}
			// u is the parent
			delta := w - val[v]
			tr.Update(in[v], in[v]+sz[v], delta)
			val[v] = w
		} else {
			v := cur[1] - 1
			res = append(res, tr.Get(in[v]))
		}
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

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	return &SegTree{
		arr: make([]int, 4*n),
		sz:  n,
	}
}

func (t *SegTree) Update(l int, r int, v int) {
	l += t.sz
	r += t.sz
	for l < r {
		if l&1 == 1 {
			t.arr[l] += v
			l++
		}

		if r&1 == 1 {
			r--
			t.arr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (t *SegTree) Get(p int) int {
	p += t.sz
	var res int
	for p > 0 {
		res += t.arr[p]
		p >>= 1
	}
	return res
}
