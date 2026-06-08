package p3952

func maxTotal(nums []int, s string) int64 {
	dp := []int{0, 0}

	const inf = 1 << 60

	for i, v := range nums {
		ndp := []int{0, 0}
		if s[i] == '0' {
			ndp[0] = max(dp[0], dp[1])
			ndp[1] = -inf
		} else {
			if i > 0 {
				ndp[0] = dp[0] + nums[i-1]
			}
			ndp[1] = max(dp[0], dp[1]) + v
		}
		copy(dp, ndp)
	}
	return int64(max(dp[0], dp[1]))
}
