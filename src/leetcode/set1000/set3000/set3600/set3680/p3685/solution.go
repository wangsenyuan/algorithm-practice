package p3685

import "sort"

func subsequenceSumAfterCapping(nums []int, k int) []bool {
	sort.Ints(nums)
	n := len(nums)

	dp := make([]bool, k+1)
	dp[0] = true
	var i int
	ans := make([]bool, n)
	for x := 1; x <= n; x++ {
		for i < n && nums[i] <= x {
			for j := k; j >= nums[i]; j-- {
				if !dp[j] {
					dp[j] = dp[j-nums[i]]
				}
			}
			i++
		}
		// 还剩余 n - i 个x
		// 前面这些数要选择
		for j, v := range dp {
			if v && (k-j)%x == 0 {
				w := (k - j) / x
				if w <= n-i {
					ans[x-1] = true
					break
				}
			}
		}
	}
	return ans
}
