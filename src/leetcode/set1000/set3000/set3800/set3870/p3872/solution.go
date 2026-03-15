package p3872

func longestArithmetic(nums []int) int {
	// 如果替换i
	n := len(nums)

	dp := make([]int, n)

	var best int
	for i := range n {
		dp[i] = 1
		if i > 0 {
			dp[i] = 2
		}
		if i >= 2 && nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
		if i > 0 {
			best = max(best, dp[i-1]+1)
		}

		best = max(best, dp[i])
	}

	fp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i+1 < n {
			best = max(best, fp[i+1]+1)
		}
		// 如果把当前给替换掉
		if i+1 < n && i > 0 && (nums[i-1]+nums[i+1])%2 == 0 {
			tmp := (nums[i-1] + nums[i+1]) / 2
			l := 1
			if i >= 2 && tmp-nums[i-1] == nums[i-1]-nums[i-2] {
				l = dp[i-1]
			}
			r := 1
			if i+2 < n && nums[i+1]-tmp == nums[i+2]-nums[i+1] {
				r = fp[i+1]
			}
			best = max(best, l+r+1)
		}
		fp[i] = 1
		if i+1 < n {
			fp[i] = 2
		}
		if i+2 < n && nums[i+1]-nums[i] == nums[i+2]-nums[i+1] {
			fp[i] = fp[i+1] + 1
		}
	}

	return best
}
