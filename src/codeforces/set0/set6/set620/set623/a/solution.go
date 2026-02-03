package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res string) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) string {
	// 先处理那些没有连接的边
	adj := make([][]int, n)
	for i := range n {
		adj[i] = make([]int, n)
		for j := range n {
			adj[i][j] = 1
		}
		adj[i][i] = 0
	}

	deg := make([]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u][v] = 0
		adj[v][u] = 0
		deg[u]++
		deg[v]++
	}

	color := make([]int, n)
	for i := range n {
		color[i] = -1
	}

	var dfs func(u int, c int) bool

	dfs = func(u int, c int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		for v := range n {
			if adj[u][v] == 1 {
				if !dfs(v, 1^c) {
					return false
				}
			}
		}
		return true
	}
	// deg[u] == n - 1 的，肯定是b（可能不存在这样的点）

	a := -1
	for i := range n {
		if deg[i] < n-1 {
			a = i
			break
		}
	}
	if a == -1 {
		// 所有的deg = n - 1, 全部设置成b
		return strings.Repeat("b", n)
	}
	// a != -1
	if !dfs(a, 0) {
		return ""
	}
	// 颜色为0和1的cnt需要知道
	cnt := make([]int, 2)
	for i := range n {
		if deg[i] < n-1 && color[i] < 0 {
			return ""
		}
		// 要么 deg[i] == n - 1, 要么 color[i] >= 0
		if color[i] >= 0 {
			cnt[color[i]]++
		}
	}

	buf := make([]byte, n)
	for i := range n {
		buf[i] = 'b'
		if color[i] >= 0 {
			if deg[i] != n-1-cnt[1^color[i]] {
				return ""
			}
			if color[i] == 0 {
				buf[i] = 'a'
			} else {
				buf[i] = 'c'
			}
		}
	}

	return string(buf)
}
