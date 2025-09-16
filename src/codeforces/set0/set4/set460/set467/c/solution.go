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
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k, m)
}

const inf = 1 << 60

func solve(a []int, k int, m int) int {
	n := len(a)
	// k * m <= n
	b := make([]int, n)
	// b[i] = sum(a[i-m+1:i])
	var sum int
	for i := range n {
		sum += a[i]
		if i-m >= 0 {
			sum -= a[i-m]
		}
		b[i] = -inf
		if i >= m-1 {
			b[i] = sum
		}
	}
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, k+1)
		for j := range k + 1 {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		if i >= m-1 {
			for j := k; j > 0; j-- {
				dp[i+1][j] = max(dp[i][j], dp[i+1-m][j-1]+b[i])
			}
		}
		dp[i+1][0] = 0
	}
	return dp[n][k]
}
