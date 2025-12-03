package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var c, d, n, m, k int
	fmt.Fscan(reader, &c, &d, &n, &m, &k)
	return solve(c, d, n, m, k)
}

func solve(c int, d int, n int, m int, k int) int {
	n1 := n*m - k
	if n1 <= 0 {
		return 0
	}
	// dp[i]表示产生i个名额需要最少的题目数
	dp := make([]int, n1+1)
	dp[0] = 0

	for i := 1; i <= n1; i++ {
		dp[i] = dp[i-1] + d
		j := max(0, i-n)
		dp[i] = min(dp[i], dp[j]+c)
	}

	return dp[n1]
}
