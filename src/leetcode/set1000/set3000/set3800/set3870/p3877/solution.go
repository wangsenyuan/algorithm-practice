package p3877

import (
	"math/bits"
	"slices"
)

const inf = 1 << 60

func minRemovals(nums []int, target int) int {
	x := slices.Max(nums)
	h := bits.Len(uint(x))
	tot := 1 << h

	if target >= tot {
		return -1
	}

	dp := make([]int, tot)
	ndp := make([]int, tot)
	for s := range tot {
		dp[s] = -inf
		ndp[s] = -inf
	}
	dp[0] = 0
	for _, v := range nums {
		for s := range tot {
			ndp[s^v] = max(ndp[s^v], dp[s]+1)
		}
		for s := range tot {
			dp[s] = max(dp[s], ndp[s])
			ndp[s] = -inf
		}
	}

	if dp[target] < 0 {
		return -1
	}
	return len(nums) - dp[target]
}
