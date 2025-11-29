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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(reader, &u, &v, &w)
		edges[i] = []int{u, v, w}
	}
	return solve(n, k, edges)
}

type data struct {
	val0 int
	val1 int
	cost int
	id   int
}

func solve(n int, k int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, cur := range edges {
		u, v, w := cur[0], cur[1], cur[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	// dp[u][0] 表示，进入u子树后，不需要返回
	// dp[u][1] 表示，进入u子树后，还必须要考虑返回的情况
	dp := make([][2]int, n)

	pref := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		var arr []data
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				arr = append(arr, data{dp[v][0], dp[v][1], g.val[i], i})
			}
		}

		// 对于子树，默认它们都要返回，
		slices.SortFunc(arr, func(a data, b data) int {
			return (b.val1 + b.cost) - (a.val1 + a.cost)
		})

		// dp[u][0] = sum(arr[i].val1) + arr[0].val0 - arr[0].val1
		// dp[u][1] = sum(arr[i].val1)
		for i, cur := range arr {
			pref[i+1] = pref[i] + cur.val1 + cur.cost
		}
		for i := 0; i < len(arr); i++ {
			if i < k {
				// 这个地方不大对
				dp[u][0] = max(dp[u][0], pref[min(len(arr), k)]+arr[i].val0-arr[i].val1)
			}

			if i >= k-1 {
				dp[u][0] = max(dp[u][0], pref[k-1]+arr[i].val0+arr[i].cost)
			}
		}
		// dp[u][1] 表示还必须回去
		dp[u][1] = pref[min(k-1, len(arr))]
	}

	dfs(-1, 0)

	return max(dp[0][0], dp[0][1])
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
