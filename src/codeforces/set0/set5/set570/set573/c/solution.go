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
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) bool {
	if n <= 4 {
		return true
	}
	adj := make([][]int, n)
	deg := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		deg[u]++
		deg[v]++
	}

	// 先把叶子节点放进去, 肯定是存在的
	var que []int
	marked := make([]bool, n)
	for u := range n {
		if len(adj[u]) == 1 {
			que = append(que, u)
			marked[u] = true
		}
	}

	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		for _, u := range adj[v] {
			deg[u]--
			if len(adj[u]) == 2 && deg[u] == 1 {
				marked[u] = true
				que = append(que, u)
			}
		}
	}
	legs := make([]int, n)

	for u := range n {
		if !marked[u] {
			for _, v := range adj[u] {
				if marked[v] {
					legs[u]++
				}
			}
		}
	}

	for u := range n {
		if !marked[u] {
			var cnt int
			for _, v := range adj[u] {
				if !marked[v] && len(adj[v])-min(2, legs[v]) > 1 {
					cnt++
				}
			}
			if cnt > 2 {
				return false
			}
		}
	}

	return true
}
