package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m, L, S, T int
	fmt.Fscan(reader, &n, &m, &L, &S, &T)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, L, S, T, edges)
}

func solve(n int, L int, S int, T int, edges [][]int) []int {
	adj := make([][][]int, n)
	for _, e := range edges {
		u, v, c := e[0]-1, e[1]-1, e[2]
		adj[u] = append(adj[u], []int{v, c})
	}

	vis := make([][]map[int]bool, n)
	for i := range n {
		vis[i] = make([]map[int]bool, L+1)
		for j := range L + 1 {
			vis[i][j] = make(map[int]bool)
		}
	}

	var res []int

	var dfs func(u int, cnt int, sum int)
	dfs = func(u int, cnt int, sum int) {
		if cnt > L || sum > T {
			return
		}
		if cnt == L {
			if sum >= S {
				res = append(res, u+1)
			}
			return
		}
		if vis[u][cnt][sum] {
			return
		}
		vis[u][cnt][sum] = true

		for _, to := range adj[u] {
			v, c := to[0], to[1]
			dfs(v, cnt+1, sum+c)
		}
	}

	dfs(0, 0, 0)

	slices.Sort(res)
	res = slices.Compact(res)

	return res
}
