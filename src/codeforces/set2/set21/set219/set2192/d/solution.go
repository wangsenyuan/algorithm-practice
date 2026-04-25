package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(a, edges)
}

func solve(a []int, edges [][]int) []int {
	n := len(a)
	adj := make([][]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// dp[u] = 子树u种最大的收获
	// 把它的一个子树v, 移动到另一棵子树w的最底端
	// sum(v) * (most_dep[w] - dep[v])
	fp := make([]int, n)
	dp := make([]int, n)
	dep := make([]int, n)
	most_dep := make([]int, n)
	sum := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sum[u] = a[u]
		most_dep[u] = dep[u]

		best := []int{-1, -1}

		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				sum[u] += sum[v]
				fp[u] += fp[v] + sum[v]
				most_dep[u] = max(most_dep[u], most_dep[v])
				if best[0] < 0 || most_dep[v] >= most_dep[best[0]] {
					best[1] = best[0]
					best[0] = v
				} else if best[1] < 0 || most_dep[v] >= most_dep[best[1]] {
					best[1] = v
				}
				dp[u] = max(dp[u], dp[v])
			}
		}
		if best[1] < 0 {
			// 只有最多一个子树
			return
		}
		for _, v := range adj[u] {
			if p != v {
				if v != best[0] {
					tmp := sum[v] * (most_dep[best[0]] - dep[v] + 1)
					dp[u] = max(dp[u], tmp)
				} else {
					tmp := sum[v] * (most_dep[best[1]] - dep[v] + 1)
					dp[u] = max(dp[u], tmp)
				}
			}
		}
	}

	dfs(-1, 0)

	for i := range n {
		fp[i] += dp[i]
	}

	return fp
}
