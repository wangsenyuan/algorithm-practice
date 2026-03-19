package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) float64 {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var res float64

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		res += 1.0 / float64(d+1)
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v, d+1)
			}
		}
	}
	dfs(-1, 0, 0)
	return res
}
