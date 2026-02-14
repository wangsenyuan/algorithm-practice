package p3839

func rob(nums []int, colors []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]
	if colors[0] != colors[1] {
		dp[1] = max(dp[1], dp[0]+nums[1])
	}
	for i := 2; i < n; i++ {
		cur := dp[i-2] + nums[i]
		if colors[i] != colors[i-1] {
			cur = max(cur, dp[i-1]+nums[i])
		}
		dp[i] = cur
		dp[i-1] = max(dp[i-1], dp[i-2])
	}
	return int64(max(dp[n-2], dp[n-1]))
}
