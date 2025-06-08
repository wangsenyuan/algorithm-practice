package p3573

const inf = 1 << 60

func maximumProfit(prices []int, k int) int64 {
	// dp[i][j][0] 表示在地i天前
	n := len(prices)
	dp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, 3)
		for j := range 3 {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := k; j >= 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i-1], dp[j][2]-prices[i-1])

			if j > 0 {
				// normal buy
				dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i-1])
				// short buy
				dp[j][2] = max(dp[j][2], dp[j-1][0]+prices[i-1])
			}
		}
	}

	var ans int
	for j := 1; j <= k; j++ {
		ans = max(ans, dp[j][0])
	}
	return int64(ans)
}
