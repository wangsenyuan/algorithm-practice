package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, res := drive(reader)
	fmt.Println(best)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (n int, edges [][]int, best int, res string) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	best, res = solve(n, edges)
	return
}

const inf = 1 << 60

func solve(n int, edges [][]int) (best int, res string) {
	adj := make([][]int, n)
	for i, e := range edges {
		v := e[1] - 1
		adj[v] = append(adj[v], i)
	}

	color := make([]int, n)
	dp := make([]int, n)
	for i := range n {
		color[i] = -1
		dp[i] = -1
	}

	var que []int
	// n-1颜色无所谓
	que = append(que, n-1)
	color[n-1] = 0
	dp[n-1] = 0

	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for _, i := range adj[u] {
			v := (edges[i][0] - 1) ^ (edges[i][1] - 1) ^ u
			w := edges[i][2]
			if color[v] == -1 {
				// 第一次访问到它，堵住
				color[v] = w ^ 1
				continue
			}
			if w == color[v] && dp[v] == -1 {
				// 颜色相同，入队
				dp[v] = dp[u] + 1
				que = append(que, v)
			}
		}
	}
	best = dp[0]
	for i := range n {
		if color[i] == -1 {
			color[i] = 0
		}
	}
	res = fmt.Sprintf("%v", color)
	res = res[1 : len(res)-1]
	res = strings.ReplaceAll(res, " ", "")
	return
}
