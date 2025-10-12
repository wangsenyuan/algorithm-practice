package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0])
	fmt.Println(res[1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	bad := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &bad[i])
	}
	return solve(n, edges, bad)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int, bad []int) []int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	cnt := make([]int, n)
	dp := make([]int, n)
	fp := make([]int, n)
	best := make([][]pair, n)

	marked := make([]bool, n)
	for _, v := range bad {
		marked[v-1] = true
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		// cnt[u]只表示子树中是否有bad men
		first := pair{0, -1}
		second := pair{0, -1}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				if marked[v] || cnt[v] > 0 {
					cnt[u]++
					fp[u] += fp[v] + 2
					// dp[u] = fp[u] - (fp[v] + 2) + dp[v] + 1
					//      = fp[u] + dp[v] - fp[v] - 1
					if dp[v]-fp[v]-1 < first.first {
						second = first
						first = pair{dp[v] - fp[v] - 1, v}
					} else if dp[v]-fp[v]-1 < second.first {
						second = pair{dp[v] - fp[v] - 1, v}
					}
				}
			}
		}

		best[u] = []pair{first, second}
		if cnt[u] > 0 {
			dp[u] = fp[u] + first.first
		}
	}

	dfs(-1, 0)

	var dfs2 func(p int, u int, x int, y int, c int)

	dfs2 = func(p int, u int, x int, y int, c int) {
		if p >= 0 && (marked[p] || c > 0) {
			cnt[u]++
			if x-y-1 < best[u][0].first {
				best[u][1] = best[u][0]
				best[u][0] = pair{x - y - 1, p}
			} else if x-y-1 < best[u][1].first {
				best[u][1] = pair{x - y - 1, p}
			}
			fp[u] += y + 2
			dp[u] = fp[u] + best[u][0].first
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			ny := fp[u]
			nc := cnt[u]
			if marked[v] || cnt[v] > 0 {
				nc--
				ny -= 2 + fp[v]
			}
			if v == best[u][0].second {
				nx := ny + best[u][1].first
				dfs2(u, v, nx, ny, nc)
			} else {
				nx := ny + best[u][0].first
				dfs2(u, v, nx, ny, nc)
			}
		}
	}
	dfs2(-1, 0, 0, 0, 0)

	ans := []int{0, dp[0]}

	for i := 1; i < n; i++ {
		if dp[i] < ans[1] {
			ans[0] = i
			ans[1] = dp[i]
		}
	}
	ans[0]++
	return ans
}

const inf = 1 << 60

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
