package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, x, y, edges)
}

func solve(n int, x int, y int, edges [][]int) int {
	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	if x >= y {
		w := slices.Max(deg)
		res := (n - 1) * y
		if w == n-1 {
			res += x - y
		}
		return res
	}
	// x < y, 尽量少的使用y
	// dp[u][0] 表示u必须和父节点通过x连接时的最优解
	// dp[u][1] 表示u子树不用和父节点连接（但是需要增加一个额外的y和外部连接）
	// dp[u][0] = sum(dp[c][1]) + c + dp[v][0] - dp[v][1] - 1
	// dp[u][1] = dp[v1][0] + dp[v2][0] + sum(dp[c][1] - ...) + c - 2

	dp := make([][2]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		var sum int
		best := []int{n, n}
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				dfs(u, v)
				sum += dp[v][1]
				tmp := dp[v][0] - dp[v][1]
				if tmp < best[0] {
					best[1] = best[0]
					best[0] = tmp
				} else if tmp < best[1] {
					best[1] = tmp
				}
			}
		}

		if cnt == 0 {
			return
		}
		dp[u][0] = sum + best[0] + cnt - 1

		if cnt == 1 {
			dp[u][1] = sum + best[0]
		} else {
			// cnt > 1
			dp[u][1] = sum + best[0] + best[1] + cnt - 2
		}
	}

	dfs(-1, 0)

	w := dp[0][1]
	res := w*y + (n-1-w)*x

	return res
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
