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
	var n int
	fmt.Fscan(reader, &n)
	color := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &color[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(color, edges)
}

func solve(color []int, edges [][]int) int {
	n := len(color)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		sz[u] = color[u]
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}
	dfs(-1, 0)

	if sz[0] == 0 || sz[0] == n {
		return 0
	}

	w := sz[0]

	var res int

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		for _, v := range adj[u] {
			if p != v && sz[v] > 0 {
				if sz[v] < w {
					res++
				}
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)
	m := res + 1

	return m - w
}
