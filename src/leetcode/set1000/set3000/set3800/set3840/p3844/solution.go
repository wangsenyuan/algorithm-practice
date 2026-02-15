package p3844

func almostPalindromic(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}

	var best int

	for r := range n {
		dp[r][r] = 0
		for l := r - 1; l >= 0; l-- {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1]
			} else {
				dp[l][r] = min(dp[l+1][r], dp[l][r-1]) + 1
			}
			if dp[l][r] < 2 {
				best = max(best, r-l+1)
			}
		}
	}
	return best
}
