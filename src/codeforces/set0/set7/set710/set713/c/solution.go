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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	n := len(a)
	for i := range n {
		a[i] -= i
	}
	b := slices.Clone(a)
	slices.Sort(b)
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = inf
		}
	}
	dp[1][0] = inf
	for i := 1; i <= n; i++ {
		dp[1][i] = min(dp[1][i-1], abs(a[0]-b[i-1]))
	}
	for i := 2; i <= n; i++ {
		dp[i][0] = inf
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]+abs(a[i-1]-b[j-1]))
		}
	}

	return dp[n][n]
}

func abs(num int) int {
	return max(num, -num)
}
