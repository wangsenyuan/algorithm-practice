package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	swaps := make([][]int, m)
	for i := range m {
		swaps[i] = make([]int, 2)
		fmt.Fscan(reader, &swaps[i][0], &swaps[i][1])
	}
	return solve(n, k, swaps)
}

const inf = 1 << 60

func solve(n int, k int, swaps [][]int) []int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = inf
	}

	dp[k-1] = 0
	for _, cur := range swaps {
		a, b := cur[0]-1, cur[1]-1
		w := dp[a]
		v := dp[b]
		dp[a] = min(w+1, v)
		dp[b] = min(w, v+1)
	}
	for i := range dp {
		if dp[i] == inf {
			dp[i] = -1
		}
	}
	return dp
}
