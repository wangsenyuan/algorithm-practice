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
	var n, k int
	fmt.Fscan(reader, &n, &k)

	universities := make([]int, 2*k)
	for i := range 2 * k {
		fmt.Fscan(reader, &universities[i])
	}

	roads := make([][]int, n-1)
	for i := range n - 1 {
		roads[i] = make([]int, 2)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1])
	}

	return solve(n, k, roads, universities)
}

func solve(n int, k int, roads [][]int, universities []int) int {
	k *= 2
	adj := make([][]int, n)
	for _, cur := range roads {
		u, v := cur[0], cur[1]
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	sz := make([]int, n)

	for _, v := range universities {
		v--
		sz[v] = 1
	}

	var res int

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for _, v := range adj[u] {
			if p != v {
				dfs(u, v)
				if sz[v] > 0 && sz[v] < k {
					res += min(sz[v], k-sz[v])
				}
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	return res
}
