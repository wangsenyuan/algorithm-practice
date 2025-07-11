package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, k, edges)
}

func solve(n int, k int, edges [][]int) int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	// dp[u][0] 表示在节点u上开始起跳的最优解
	// dp[u][1] 表示在节点u的直接子节点v上，开始起跳的最优解
	sz := make([]int, n)
	// 这个需要和u距离为k的节点的sz之和
	// 要怎么更新呢？
	dp := make([][]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dp[u] = make([]int, k)
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)

				for j := 0; j < k; j++ {
					dp[u][(j+1)%k] += dp[v][j]
				}
				dp[u][0] += sz[v]
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)
	var ans int

	var dfs2 func(p int, u int, d []int)
	dfs2 = func(p int, u int, d []int) {
		if p >= 0 {
			for j := 0; j < k; j++ {
				dp[u][(j+1)%k] += d[j]
			}
			dp[u][0] += n - sz[u]
		}

		sz[u] = n

		ans += dp[u][0]

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				// 先要把v的贡献给消除掉
				nd := slices.Clone(dp[u])

				for j := 0; j < k; j++ {
					nd[(j+1)%k] -= dp[v][j]
				}
				nd[0] -= sz[v]

				dfs2(u, v, nd)
			}
		}
	}

	dfs2(-1, 0, nil)

	return ans / 2
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
