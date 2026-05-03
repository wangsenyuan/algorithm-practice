package p3919

func minCost(nums []int, queries [][]int) []int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1]
		if i == 0 || nums[i+1]-nums[i] < nums[i]-nums[i-1] {
			dp[i]++
		} else {
			dp[i] += nums[i+1] - nums[i]
		}
	}

	fp := make([]int, n)
	for i := 1; i < n; i++ {
		fp[i] = fp[i-1]
		if i == n-1 || nums[i]-nums[i-1] <= nums[i+1]-nums[i] {
			fp[i]++
		} else {
			fp[i] += nums[i] - nums[i-1]
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l < r {
			ans[i] = dp[l] - dp[r]
		} else if l > r {
			ans[i] = fp[l] - fp[r]
		}
	}
	return ans
}
