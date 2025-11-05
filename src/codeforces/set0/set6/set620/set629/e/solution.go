package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	ans := drive(reader)
	for _, x := range ans {
		fmt.Fprintf(writer, "%.10f\n", x)
	}
}

func drive(reader *bufio.Reader) []float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	friends := make([][]int, m)
	for i := range m {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	return solve(n, edges, friends)
}

func solve(n int, edges [][]int, friends [][]int) []float64 {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dep := make([]int, n)

	h := bits.Len(uint(n))
	fa := make([][]int, n)

	sz := make([]int, n)
	dp := make([]int, n)
	dp2 := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {

		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
				dp[u] += dp[v] + sz[v]
			}
		}
		// dp[u]表示如果在u得某个子节点v，添加一条边后， 形成一个cycle, v节点对u处得贡献之和
		dp2[u] = dp[u]
	}

	dfs(0, 0)

	var dfs2 func(p int, u int, ph int)

	dfs2 = func(p int, u int, ph int) {
		dp2[u] += ph
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dp2[u] - (dp[v] + sz[v])
				dfs2(u, v, tmp+n-sz[v])
			}
		}
	}

	dfs2(-1, 0, 0)

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

	kth := func(u int, k int) int {
		for i := h - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = fa[u][i]
			}
		}
		return u
	}

	// 这个不对，这样算出来的是最好的结果
	find := func(u int, v int) float64 {
		if dep[u] > dep[v] {
			u, v = v, u
		}
		// 假设 u ~ p ~ v
		// 只有当多出来的边，添加到(u, v)的外部的时候，才能形成一个cycle
		// 假设是  a ~ u ~ p ~ v ~ b
		// 那么 a b 的贡献  = dist(u, v) + dist(a, u) + dist(v, b) + 1
		//     a, c 的贡献  = dist(u, v) + dist(a, u) + dist(v, c) + 1
		//  ... (dist(u, v) + 1) * sz[v]  + dp[v] + sz[v] * dist(a, u)
		// ... (dist(u, v) + 1) * sz[v] + dp[v] + sz[v] * dist(a1, u)
		// ...
		// (dist(u, v) + 1) * sz[v] * sz[u] + dp[v] * sz[u] + sz[v] * dp[u]

		p := lca(u, v)

		if u == p {
			w := kth(v, dep[v]-dep[u]-1)
			res := (dep[v] - dep[u] + 1) * sz[v] * (n - sz[w])
			res += dp[v] * (n - sz[w])
			res += (dp2[u] - dp[w] - sz[w]) * sz[v]
			cnt := sz[v] * (n - sz[w])
			return float64(res) / float64(cnt)
		}

		res := (dep[u] + dep[v] - 2*dep[p] + 1) * sz[v] * sz[u]
		res += dp[v] * sz[u]
		res += dp[u] * sz[v]

		cnt := sz[v] * sz[u]
		return float64(res) / float64(cnt)
	}

	ans := make([]float64, len(friends))

	for i, cur := range friends {
		u, v := cur[0]-1, cur[1]-1
		ans[i] = find(u, v)
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
