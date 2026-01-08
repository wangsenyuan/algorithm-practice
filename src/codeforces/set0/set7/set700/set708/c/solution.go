package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

type data struct {
	c int
	w int
}

func solve(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)
	dp := make([][]data, n)

	update := func(cur []data, d data) {
		if d.w >= cur[0].w {
			cur[1] = cur[0]
			cur[0] = d
		} else if d.w >= cur[1].w {
			cur[1] = d
		}
	}

	fa := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = 1
		dp[u] = make([]data, 2)
		fa[u] = p
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
				// 只能记录 sz[v] <= n/2的子树
				if sz[v] <= n/2 {
					update(dp[u], data{v, sz[v]})
				} else {
					update(dp[u], data{v, dp[v][0].w})
				}
			}
		}
	}

	dfs(-1, 0)

	var dfs2 func(p int, u int, w int)
	dfs2 = func(p int, u int, w int) {
		if p >= 0 {
			if n-sz[u] <= n/2 {
				update(dp[u], data{p, n - sz[u]})
			} else {
				update(dp[u], data{p, w})
			}
		}

		for _, v := range adj[u] {
			if p != v {
				if dp[u][0].c == v {
					dfs2(u, v, dp[u][1].w)
				} else {
					dfs2(u, v, dp[u][0].w)
				}
			}
		}
	}

	dfs2(-1, 0, 0)

	check := func(u int) int {
		big := data{-1, 0}
		for _, v := range adj[u] {
			tmp := sz[v]
			if v == fa[u] {
				tmp = n - sz[u]
			}
			if tmp > big.w {
				big = data{v, tmp}
			}
		}

		if big.w <= n/2 {
			return 1
		}
		// big.w > n / 2
		v := big.c
		next := dp[v][0]
		if next.c == u {
			next = dp[v][1]
		}
		// next.w <= n / 2
		if next.w > 0 && big.w-next.w <= n/2 {
			return 1
		}
		return 0
	}

	ans := make([]int, n)

	for u := range n {
		ans[u] = check(u)
	}

	return ans
}
