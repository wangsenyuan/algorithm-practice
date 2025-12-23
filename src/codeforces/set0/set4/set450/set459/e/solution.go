package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	slices.SortFunc(edges, func(x []int, y []int) int {
		return cmp.Or(x[2]-y[2], x[0]-y[0], x[1]-y[1])
	})

	// dp[u] = 终点在u的最长的路径
	dp := make([]int, n)

	m := len(edges)

	fp := make([]int, n)

	for i := 0; i < m; {
		j := i
		for i < m && edges[j][2] == edges[i][2] {
			u, v := edges[i][0]-1, edges[i][1]-1
			fp[v] = max(fp[v], dp[u]+1)
			i++
		}
		for i1 := j; i1 < i; i1++ {
			v := edges[i1][1] - 1
			dp[v] = max(dp[v], fp[v])
			fp[v] = -1
		}
	}

	return slices.Max(dp)
}
