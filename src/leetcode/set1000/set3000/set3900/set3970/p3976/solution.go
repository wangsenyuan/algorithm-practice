package p3976

func maxSubarraySum(nums []int, k int) int64 {
	// dp[0] 表示到i为止,没有进行操作, 最大和
	// dp[1] 表示到i为止,使用 *k
	// dp[2] 表示到i为止,使用/k
	// dp[3] 表示已经使用了一次操作后的,最大和

	dp := make([]int, 4)
	best := -(1 << 60)

	for _, v := range nums {
		s0 := dp[0] + v
		// 可以从当前位置开始 *k
		s1 := max(dp[0], dp[1]) + v*k
		s2 := max(dp[0], dp[2]) + v/k
		s3 := max(dp[3], dp[1], dp[2]) + v

		best = max(best, s0, s1, s2, s3)

		dp[0] = max(0, s0)
		dp[1] = max(0, s1)
		dp[2] = max(0, s2)
		dp[3] = max(0, s3)
	}

	return int64(best)
}
