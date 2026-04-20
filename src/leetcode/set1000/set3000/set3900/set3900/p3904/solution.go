package p3904

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = nums[i]
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}
	res := -1
	suf := nums[n-1]
	for i := n - 1; i >= 0; i-- {
		suf = min(suf, nums[i])
		if dp[i]-suf <= k {
			res = i
		}
	}
	return res
}
