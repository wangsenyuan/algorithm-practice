package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := NewGraph(n, n*2)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	sz := make([]int, n)

	var tot int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if p != v {
				dfs(u, v)
				// 这条边的贡献 = 两边节点的数量的乘积
				tot += edges[j][2] * sz[v] * (n - sz[v])
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	dp := make([]int, n)

	var dfs2 func(p int, u int, blockedEdge int)

	dfs2 = func(p int, u int, blockedEdge int) {
		dp[u] = 0
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if p == v || j == blockedEdge {
				continue
			}
			dfs2(u, v, blockedEdge)
			sz[u] += sz[v]
			dp[u] += dp[v] + edges[j][2]*sz[v]
		}
	}

	var dfs3 func(p int, u int, blockedEdge int, m int, fp int, res *int)

	dfs3 = func(p int, u int, blockedEdge int, m int, fp int, res *int) {
		// dp[u]越小越好
		*res = min(*res, dp[u]+fp)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if p == v || j == blockedEdge {
				continue
			}
			nfp := fp + edges[j][2]*(m-sz[v]) + dp[u] - dp[v] - edges[j][2]*sz[v]
			dfs3(u, v, blockedEdge, m, nfp, res)
		}
	}

	find := func(r int, blockedEdge int) int {
		dfs2(-1, r, blockedEdge)
		cur := dp[r]
		res := cur
		dfs3(-1, r, blockedEdge, sz[r], 0, &res)
		return cur - res
	}

	var best int

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		// 如果把这条边替换掉
		// d1是在u子树中找到的最优的点，得到的最优的节省的距离
		d1 := find(u, i)
		d2 := find(v, i)

		tmp := d1*sz[v] + d2*sz[u]

		best = max(best, tmp)
	}

	return tot - best
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
