package p3738

func longestSubarray(nums []int) int {
	n := len(nums)

	dp := make([]int, n)

	var best int
	for i := range n {
		if i == 0 || nums[i] < nums[i-1] {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
		}
		best = max(best, dp[i]+check(i+1 < n))
	}

	var fp int
	for i := n - 1; i >= 0; i-- {
		// 修改i
		if i > 0 && i+1 < n && nums[i-1] <= nums[i+1] {
			best = max(best, fp+dp[i-1]+1)
		}
		if i == n-1 || nums[i] > nums[i+1] {
			fp = 1
		} else {
			fp++
		}
		best = max(best, fp+check(i > 0))
	}

	return best
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
