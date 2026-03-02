package p3857

const inf = 1 << 60

func minCost(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = inf
		for j := 1; j < i; j++ {
			dp[i] = min(dp[i], dp[j]+dp[i-j]+j*(i-j))
		}
	}

	return dp[n]
}
