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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, k, edges)
}

func solve(n int, k int, edges [][]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)

	var res int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for _, v := range adj[u] {
			if v != p {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
		sz[u]++

		if sz[u] >= k {
			// 如果父节点在p中
			res += n - sz[u]
		}

		for _, v := range adj[u] {
			if p != v && n-sz[v] >= k {
				// 如果父节点在v子树中
				res += sz[v]
			}
		}
		// u自己也可以作为lca
		res++
	}

	dfs(-1, 0)

	return res
}
