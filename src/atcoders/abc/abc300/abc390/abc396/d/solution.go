package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		var u, v int
		var w int
		fmt.Fscan(reader, &u, &v, &w)
		edges[i] = []int{u, v, w}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {

	adj := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], i)
		adj[v] = append(adj[v], i)
	}

	ans := 1 << 61

	var dfs func(u int, mask int, sum int)

	dfs = func(u int, mask int, sum int) {
		if u == n-1 {
			ans = min(ans, sum)
			return
		}
		for _, i := range adj[u] {
			v := (edges[i][0] ^ edges[i][1] ^ (u + 1)) - 1
			w := edges[i][2]
			if (mask>>v)&1 == 0 {
				dfs(v, mask|(1<<v), sum^w)
			}
		}
	}

	dfs(0, 1, 0)

	return ans
}
