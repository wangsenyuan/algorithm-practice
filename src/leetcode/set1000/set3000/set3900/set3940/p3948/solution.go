package p3948

func maximumMEX(nums []int) []int {
	n := len(nums)
	// 从后往前处理一遍
	vis := make([]bool, n+2)
	dp := make([]int, n+1)
	var mex int
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= n+1 {
			vis[nums[i]] = true
		}
		for vis[mex] {
			mex++
		}
		dp[i] = mex
	}
	// 找到最短的左边的序列 = dp[i]
	var res []int
	mark := make([]int, n+2)
	var cur int
	marker := 1
	for i, v := range nums {
		if v <= n+1 {
			mark[v] = marker
		}
		for mark[cur] == marker {
			cur++
		}
		if cur == mex {
			// 需要把vis更更新掉
			res = append(res, mex)
			mex = dp[i+1]
			cur = 0
			marker = i + 2
		}
	}
	return res
}
