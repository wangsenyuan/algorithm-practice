package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.10f\n", res)
}

func process(reader *bufio.Reader) float64 {
	var n, t int
	var p float64
	fmt.Fscan(reader, &n, &p, &t)
	return solve(n, p, t)
}

func solve(n int, p float64, t int) float64 {
	// dp[i][t] = 在第t秒有i个人进入了电梯的概率
	// dp[i+1][t] += dp[i][t-1] * p
	// dp[i+1][t] += dp[i+1][t-1] * (1 - p)

	dp := make([]float64, n+1)
	dp[0] = 1
	for range t {
		// dp[n] 不能变了
		for i := n - 1; i >= 0; i-- {
			dp[i+1] += dp[i] * p
			dp[i] *= (1 - p)
		}
	}
	var res float64
	for i := 0; i <= n; i++ {
		res += dp[i] * float64(i)
	}
	return res
}
