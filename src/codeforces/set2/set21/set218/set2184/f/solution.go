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

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []string {
	var tc int
	fmt.Fscan(reader, &tc)
	ans := make([]string, tc)
	for i := range tc {
		var n int
		fmt.Fscan(reader, &n)
		edges := make([][]int, n-1)
		for j := range n - 1 {
			edges[j] = make([]int, 2)
			fmt.Fscan(reader, &edges[j][0], &edges[j][1])
		}
		ans[i] = solve(n, edges)
	}
	return ans
}

func solve(n int, edges [][]int) string {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dp := make([][]bool, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		var prev []bool
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				if len(prev) == 0 {
					prev = dp[v]
				} else {
					cur := make([]bool, 3)
					for i, x := range prev {
						if x {
							for j, y := range dp[v] {
								if y {
									cur[(i+j)%3] = true
								}
							}
						}
					}
					prev = cur
				}
			}
		}

		if len(prev) == 0 {
			prev = make([]bool, 3)
		}
		prev[1] = true
		dp[u] = prev
	}

	dfs(0, 0)

	if dp[0][0] {
		return "Yes"
	}
	return "No"
}
