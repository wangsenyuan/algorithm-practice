package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

const inf = 1 << 40

func solve(n int, edges [][]int) float64 {
	adj := make([][]int, n)
	for _, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(p int, u int) float64

	dfs = func(p int, u int) float64 {

		var sum float64
		for _, v := range adj[u] {
			if p != v {
				sum += dfs(u, v) + 1
			}
		}

		if sum == 0 {
			return 0
		}

		k := len(adj[u])
		if p != -1 {
			k--
		}
		return sum / float64(k)
	}

	return dfs(-1, 0)
}
