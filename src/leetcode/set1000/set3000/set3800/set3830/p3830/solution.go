package p3830

func longestAlternating(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n)

	var res int

	for i := 0; i < n; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
		if i > 0 {
			if nums[i] < nums[i-1] {
				dp[i][0] = dp[i-1][1] + 1
			}
			if nums[i] > nums[i-1] {
				dp[i][1] = dp[i-1][0] + 1
			}
		}
		res = max(res, dp[i][0], dp[i][1])
	}

	fp := make([]int, 2)
	for i := n - 1; i > 0; i-- {
		// 如果删除字符i
		if i+1 < n {
			if nums[i-1] < nums[i+1] {
				res = max(res, dp[i-1][0]+fp[1])
			}
			if nums[i-1] > nums[i+1] {
				res = max(res, dp[i-1][1]+fp[0])
			}
		}
		var x, y int
		if i == n-1 || nums[i] <= nums[i+1] {
			x = 1
		} else {
			x = fp[0] + 1
		}
		if i == n-1 || nums[i] >= nums[i+1] {
			y = 1
		} else {
			y = fp[1] + 1
		}
		fp[0] = y
		fp[1] = x
	}

	return res
}
