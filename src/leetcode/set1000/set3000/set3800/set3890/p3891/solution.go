package p3891

func minIncrease(nums []int) int64 {
	n := len(nums)
	play := func(s int) int {
		var res int
		for i := s; i+1 < n; i += 2 {
			res += max(0, max(nums[i-1], nums[i+1])+1-nums[i])
		}
		return res
	}

	if n%2 == 1 {
		return int64(play(1))
	}

	// 只在1,3,5,7
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		if i&1 == 1 && i < n-1 {
			dp[i] += max(0, max(nums[i-1], nums[i+1])+1-nums[i])
		}
	}
	res := dp[n-1]

	var fp int
	for i := n - 2; i >= 2; i-- {
		if i&1 == 0 {
			fp += max(0, max(nums[i-1], nums[i+1])+1-nums[i])
			if i >= 3 {
				res = min(res, dp[i-3]+fp)
			}
		}

	}

	return int64(min(res, fp))
}
