package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	type data struct {
		leaf int
		to   int
		dist int
	}

	dp := make([][]data, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		dp[u] = append(dp[u], data{u, u, 0})
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				tmp := data{dp[v][0].leaf, v, dp[v][0].dist + 1}
				for i, cur := range dp[u] {
					if tmp.dist > cur.dist || tmp.dist == cur.dist && tmp.leaf > cur.leaf {
						dp[u][i], tmp = tmp, cur
					}
				}

				if len(dp[u]) < 2 {
					dp[u] = append(dp[u], tmp)
				}
			}
		}
	}

	dfs(-1, 0)

	ans := make([]int, n)

	var dfs2 func(p int, u int, from data)
	dfs2 = func(p int, u int, from data) {
		for i, cur := range dp[u] {
			if from.dist > cur.dist || from.dist == cur.dist && from.leaf > cur.leaf {
				dp[u][i], from = from, cur
			}
		}

		if len(dp[u]) < 2 {
			dp[u] = append(dp[u], from)
		}

		ans[u] = dp[u][0].leaf + 1

		for _, v := range adj[u] {
			if p != v {
				tmp := dp[u][0]
				if tmp.to == v {
					tmp = dp[u][1]
				}
				dfs2(u, v, data{tmp.leaf, u, tmp.dist + 1})
			}
		}
	}

	dfs2(-1, 0, data{-1, -1, -n})

	return ans
}
