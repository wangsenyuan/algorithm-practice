package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func drive(in *bufio.Reader, out *bufio.Writer) {
	buf := make([]byte, 4096)
	var _i, _n int
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	tc := rd()

	for range tc {
		n := rd()
		m := rd()
		c := rd()
		edges := make([][]int, m)
		for i := range m {
			edges[i] = []int{rd(), rd()}
		}
		res := solve(n, edges, c)
		fmt.Fprintln(out, res)
	}
}

const inf = 1 << 62

type Pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int, c int) int {
	// c is useless in this problem
	g := NewGraph(n, len(edges)*2)

	for _, e := range edges {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	// 需要将那些strong component缩放成一个点
	stack := make([]int, n)
	var top int
	low := make([]int, n)
	dis := make([]int, n)
	vis := make([]bool, n)

	belong := make([]int, n)
	var comp []int

	var time int

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		dis[u] = time
		low[u] = time
		time++
		vis[u] = true

		stack[top] = u
		top++

		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if !vis[v] {
					sz += dfs(u, v)
					low[u] = min(low[u], low[v])
				} else {
					low[u] = min(low[u], dis[v])
				}
			}
		}

		if low[u] == dis[u] {
			// 这部分是一个强连通分量
			id := len(comp)
			var cnt int
			for top > 0 {
				v := stack[top-1]
				top--
				belong[v] = id
				cnt++
				if u == v {
					break
				}
			}
			comp = append(comp, cnt)
		}

		return sz
	}
	var trs []Pair

	for i := 0; i < n; i++ {
		if !vis[i] {
			sz := dfs(-1, i)
			trs = append(trs, Pair{sz, i})
		}
	}
	m := len(comp)

	if m == 1 {
		// can't split
		return -1
	}

	// g2 is a forest for those trees
	g2 := NewGraph(m, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u = belong[u]
		v = belong[v]
		if u != v {
			g2.AddEdge(u, v)
			g2.AddEdge(v, u)
		}
	}

	f1 := NewBitSet(n + 1)
	f0 := NewBitSet(n + 1)
	f0.Set(0)

	var dfs2 func(p int, u int, k int) int
	dfs2 = func(p int, u int, k int) int {
		sz := comp[u]
		for i := g2.nodes[u]; i > 0; i = g2.next[i] {
			v := g2.to[i]
			if p != v {
				sz += dfs2(u, v, k)
			}
		}
		if p >= 0 {
			a := f0.Copy()
			a.LeftShift(sz)
			b := f0.Copy()
			b.LeftShift(k - sz)
			f1.Union(a)
			f1.Union(b)
		}

		return sz
	}

	k := len(trs)

	leftShift := func(bs *BitSet, cnt int) {
		tmp := bs.Copy()
		tmp.LeftShift(cnt)
		bs.Union(tmp)
	}

	for i := 0; i < k; i++ {
		cnt, root := trs[i].first, trs[i].second
		leftShift(f1, cnt)
		dfs2(-1, belong[root], cnt)
		leftShift(f0, cnt)
	}

	best := inf

	for i := 1; i <= n/2; i++ {
		if f1.IsSet(i) || f0.IsSet(i) {
			best = min(best, i*i+(n-i)*(n-i))
		}
	}

	// a = k-1
	return best + c*(k-1)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}

func (bs *BitSet) Count() int {
	var res int
	for i := 0; i < bs.sz; i++ {
		res += countDigit(bs.arr[i])
	}
	return res
}

func countDigit(num uint64) int {
	var res int
	if (num>>63)&1 == 1 {
		res++
	}
	tmp := int64(num & ((1 << 63) - 1))

	for tmp > 0 {
		res++
		tmp -= tmp & -tmp
	}
	return res
}

func (bs *BitSet) LeftShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := 0; u+i < bs.sz; u++ {
			bs.arr[u] = bs.arr[u+i]
		}
	} else {
		for u := 0; u+i < bs.sz; u++ {
			v := u + i
			bs.arr[u] = bs.arr[v] << uint64(j)
			if v+1 < bs.sz {
				bs.arr[u] |= bs.arr[v+1] >> uint64(64-j)
			}
		}
	}
	for u := bs.sz - i; u < bs.sz; u++ {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) RightShift(p int) {
	i, j := p/64, p%64
	if j == 0 {
		for u := bs.sz - 1; u-i >= 0; u-- {
			bs.arr[u] = bs.arr[u-i]
		}
	} else {
		for u := bs.sz - 1; u-i >= 0; u-- {
			v := u - i
			bs.arr[u] = bs.arr[v] >> uint64(j)
			if v-1 >= 0 {
				bs.arr[u] |= bs.arr[v-1] << uint64(64-j)
			}
		}
	}
	for u := i - 1; u >= 0; u-- {
		bs.arr[u] = 0
	}
}

func (bs *BitSet) Copy() *BitSet {
	res := NewBitSet(len(bs.arr) * 64)
	copy(res.arr, bs.arr)
	return res
}

func (this *BitSet) Union(that *BitSet) {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] |= that.arr[i]
	}
}

func (this *BitSet) Reset() {
	clear(this.arr)
}
