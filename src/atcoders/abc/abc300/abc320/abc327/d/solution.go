package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, m)
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(n, a, b)
}

func solve(n int, a, b []int) bool {
	adj := make([][]int, n+1)
	for i := range a {
		adj[a[i]] = append(adj[a[i]], b[i])
		adj[b[i]] = append(adj[b[i]], a[i])
	}

	color := make([]int, n+1)
	for i := range n + 1 {
		color[i] = -1
	}

	var dfs func(v int, c int) bool
	dfs = func(v int, c int) bool {
		if color[v] != -1 {
			return color[v] == c
		}
		color[v] = c
		for _, u := range adj[v] {
			if !dfs(u, c^1) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if color[i] == -1 && !dfs(i, 0) {
			return false
		}
	}
	return true
}
