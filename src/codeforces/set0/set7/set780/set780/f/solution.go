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
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {

	dp := make([][][]map[int]bool, 60)

	setup := func(d int) {
		dp[d] = make([][]map[int]bool, n)
		for i := range n {
			dp[d][i] = make([]map[int]bool, 2)
			for j := range 2 {
				dp[d][i][j] = make(map[int]bool)
			}
		}
	}

	setup(0)

	for _, e := range edges {
		u, v, t := e[0], e[1], e[2]
		u--
		v--
		dp[0][u][t][v] = true
	}

	if dp[0][0][0][0] && dp[0][0][1][0] {
		return -1
	}

	var find func(d int, x int, w int) int
	find = func(d int, x int, w int) int {
		if d < 0 {
			return 0
		}
		var res int
		for u := range dp[d][x][w] {
			res = max(res, find(d-1, u, w^1))
		}
		return res + 1<<d
	}

	for d := range 59 {
		setup(d + 1)
		for x := range n {
			for t := range 2 {
				for u := range dp[d][x][t] {
					for v := range dp[d][u][t^1] {
						dp[d+1][x][t][v] = true
					}
				}
			}
		}
		if dp[d+1][0][0][0] && dp[d+1][0][1][0] {
			return -1
		}

		if len(dp[d+1][0][0]) == 0 {
			// dp[d][0][0][0] is true
			return find(d, 0, 0)
		}
	}

	return -1
}
