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
	friends := make([][]int, m)
	for i := 0; i < m; i++ {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	return solve(n, friends)
}

func solve(n int, friends [][]int) int {
	// 对于任何一对(a, b), 找到c， 使的sum(a, b, c)最小
	adj := make([]map[int]int, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make(map[int]int)
	}

	for _, cur := range friends {
		u, v := cur[0], cur[1]
		adj[u][v]++
		adj[v][u]++
	}

	res := -1
	for _, cur := range friends {
		u, v := cur[0], cur[1]
		for w := 1; w <= n; w++ {
			if w != u && w != v && adj[u][w] > 0 && adj[w][v] > 0 {
				sum := len(adj[u]) + len(adj[v]) + len(adj[w]) - 6
				if res < 0 || sum < res {
					res = sum
				}
			}
		}
	}
	return res
}
