package p3693

const inf = 1 << 60

func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j <= min(n, i+3); j++ {
			dp[j] = min(dp[j], dp[i]+costs[j-1]+(j-i)*(j-i))
		}
	}
	return dp[n]
}
