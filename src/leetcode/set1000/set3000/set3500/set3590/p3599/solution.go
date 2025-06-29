package p3599

const inf = 1 << 30

func minXor(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, k+1)
		for j := range k + 1 {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for r := 1; r <= n; r++ {
		var sum int
		for l := r - 1; l >= 0; l-- {
			sum ^= nums[l]
			for j := 0; j+1 <= k; j++ {
				dp[r][j+1] = min(dp[r][j+1], max(dp[l][j], sum))
			}
		}
	}
	return dp[n][k]
}
