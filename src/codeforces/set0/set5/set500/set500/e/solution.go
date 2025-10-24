package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	l := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i], &l[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(n, l, p, queries)
}

const inf = 1 << 60

func solve(n int, l []int, p []int, queries [][]int) []int {
	far := NewSegTree(n+1, -1, func(a, b int) int {
		return max(a, b)
	})

	hi := NewSegTree(n+1, 0, func(a, b int) int {
		return max(a, b)
	})

	p = append(p, inf)

	g := NewGraph(n+1, n)

	for i := n - 1; i >= 0; i-- {
		hi.Update(i, p[i]+l[i])

		j := sort.Search(n+1, func(j int) bool {
			return p[j] > p[i]+l[i]
		})
		// p[i] + l[i] >= p[j]
		j--
		if j+1 == n {
			// 可以到达终点
			g.AddEdge(j+1, i, 0)
		} else if i == j {
			// push不到i+1
			g.AddEdge(j+1, i, p[j+1]-l[j]-p[j])
		} else {
			// 要比i能到达的更远
			j = far.Get(i+1, j+1)
			// 先到达j, 花费p[j+1] - l[j] 能够到达j+
			// 这个cost似乎不大对
			// 应该是 p[j+1] - l[?] where p[?] + l[?]离 p[j+1]最近
			// 也就是要找出i...j中间最大的 p[i] + l[i]
			v := hi.Get(i, j+1)
			g.AddEdge(j+1, i, p[j+1]-v)
		}
		// 如果能push i, 就可以push j
		far.Update(i, j)
	}

	h := bits.Len(uint(n + 1))
	fa := make([][]int, n+1)
	// dp[i]是从n到i的cost
	dp := make([]int, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			dp[v] = dp[u] + w
			dfs(u, v)
		}
	}

	dfs(n, n)

	find := func(x int, y int) int {
		// 找到x的祖先节点p, p >= y
		p := x
		for i := h - 1; i >= 0; i-- {
			if fa[p][i] < y {
				// fa[p][i+1] >= y
				p = fa[p][i]
			}
		}
		if fa[p][0] == y {
			p = y
		}
		return dp[x] - dp[p]
	}

	res := make([]int, len(queries))
	for i, cur := range queries {
		x, y := cur[0]-1, cur[1]-1
		res[i] = find(x, y)
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
