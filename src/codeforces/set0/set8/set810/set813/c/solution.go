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
	var n, x int
	fmt.Fscan(reader, &n, &x)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, x, edges)
}

func solve(n int, x int, edges [][]int) int {
	// 这里alice肯定一直往x所在的节点移动
	// x只能尽量的往离root远的节点移动
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	x--
	fa := make([]int, n)
	dp := make([]int, n)
	dep := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				dp[u] = max(dp[u], dp[v]+1)
			}
		}
	}

	dfs(0, 0)
	// 往x最远的那个叶子移动
	res := 2 * (dep[x] + dp[x])

	var dist int
	for u := x; u != 0; u = fa[u] {
		if dep[u] <= dist {
			break
		}
		// dep[u] > dist
		res = max(res, (dp[u]+dep[u])*2)
		dist++
	}

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
