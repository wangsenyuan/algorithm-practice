package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)

	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}

	return solve(n, k, edges)
}

const inf = 1 << 60

func solve(n int, k int, edges [][]int) int {
	// dp[l][r][u][j]表示在范围[l...r]内，现在在u位置，访问了j条边时的最优解

	// 80 * 80 * 80 * 80 = 8 * 8 * 8 * 8 * 10000
	// = 40960000
	// 确实太大了，得去掉一个维度

	g := NewGraph(n+2, len(edges))

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
	}

	dp := make([][][]int, n+2)
	ndp := make([][][]int, n+2)
	for i := range dp {
		dp[i] = make([][]int, n+2)
		ndp[i] = make([][]int, n+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+2)
			ndp[i][j] = make([]int, n+2)
			for u := range n + 2 {
				dp[i][j][u] = inf
			}
		}
	}
	for u := 1; u <= n; u++ {
		dp[0][n+1][u] = 0
	}

	for range k {
		for l := 0; l < n; l++ {
			for r := 2; r <= n+1; r++ {
				for u := l + 1; u < r; u++ {
					for i := g.nodes[u]; i > 0; i = g.next[i] {
						v := g.to[i]
						w := g.val[i]
						if l < v && v < u {
							ndp[l][u][v] = min(ndp[l][u][v], dp[l][r][u]+w)
						}
						if u < v && v < r {
							ndp[u][r][v] = min(ndp[u][r][v], dp[l][r][u]+w)
						}
					}
				}
			}
		}
		for l := 0; l <= n; l++ {
			for r := l + 1; r <= n+1; r++ {
				for u := l; u <= r; u++ {
					dp[l][r][u] = ndp[l][r][u]
					ndp[l][r][u] = inf
				}
			}
		}
	}

	ans := inf
	for l := 0; l <= n; l++ {
		for r := l + 2; r <= n+1; r++ {
			for u := l + 1; u < r; u++ {
				ans = min(ans, dp[l][r][u])
			}
		}
	}
	if ans == inf {
		return -1
	}

	return ans
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
