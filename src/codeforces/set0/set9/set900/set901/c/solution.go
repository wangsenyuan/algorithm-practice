package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	var buf bytes.Buffer
	for _, v := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(n, edges, queries)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*len(edges))
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, 1)
		g.AddEdge(v, u, 1)
	}

	dfn := make([]int, n)
	low := make([]int, n)
	stack := make([]int, n)
	var top int
	var timer int
	in_stack := make([]bool, n)
	var arr []pair

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		timer++
		dfn[u] = timer
		low[u] = timer
		stack[top] = u
		top++
		in_stack[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p {
				continue
			}
			if in_stack[v] {
				low[u] = min(low[u], dfn[v])
			} else if dfn[v] == 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			}
		}

		if low[u] == dfn[u] {
			cur := pair{u, u}
			for top > 0 {
				v := stack[top-1]
				top--
				cur.first = min(cur.first, v)
				cur.second = max(cur.second, v)
				in_stack[v] = false
				if u == v {
					break
				}
			}
			arr = append(arr, cur)
		}
	}

	for u := 0; u < n; u++ {
		if dfn[u] == 0 {
			dfs(-1, u)
		}
	}
	R := make([]int, n+2)
	for i := range n + 2 {
		R[i] = n
	}
	for _, p := range arr {
		if p.first < p.second {
			R[p.first+1] = p.second
		}
	}
	for i := n - 1; i > 0; i-- {
		R[i] = min(R[i], R[i+1])
	}

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + R[i]
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		pos := cur[1] + 1

		for l <= r {
			mid := (l + r) >> 1
			if R[mid] > cur[1] {
				pos = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		pos--
		l, r = cur[0], cur[1]
		ans[i] = (r-pos)*r + dp[pos] - dp[l-1] - (r+l)*(r-l+1)/2 + r - l + 1
	}

	return ans
}

type Graph struct {
	nodes []int
	to    []int
	next  []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	to := make([]int, e+1)
	next := make([]int, e+1)
	val := make([]int, e+1)
	return &Graph{nodes, to, next, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
