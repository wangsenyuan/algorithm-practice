package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	res := solve(n, k)
	fmt.Println(res)
}

const mod = 10000

func solve(n int, k int) int {
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[j] = (dp[j] + dp[j-1]) % mod
		}
		for j := k; j > i; j-- {
			dp[j] = (dp[j] - dp[j-i-1] + mod) % mod
		}
	}

	return dp[k]
}

func solve1(n int, k int) int {

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}

	var f func(n int, k int) int

	f = func(n int, k int) (ans int) {
		if n == 1 {
			if k == 0 {
				return 1
			}
			return 0
		}

		if dp[n][k] >= 0 {
			return dp[n][k]
		}

		defer func() {
			dp[n][k] = ans
		}()

		// j是从后当前的位置
		for j := 0; j < n && j <= k; j++ {
			ans += f(n-1, k-j)
		}

		return ans % mod
	}

	return f(n, k)
}
