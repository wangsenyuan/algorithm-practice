package p3670

import (
	"math/bits"
	"slices"
)

func maxProduct(nums []int) int64 {
	// n := len(nums)
	x := slices.Max(nums)
	h := bits.Len(uint(x))

	dp := make([]int, 1<<h)

	marked := make([]bool, 1<<h)
	for _, num := range nums {
		marked[num] = true
	}

	// 假设是4，那么他可以和 3,2,1 进行乘法运算
	all := 1<<h - 1
	var best int
	for mask := 1; mask < 1<<h; mask++ {
		if marked[mask] {
			dp[mask] = max(dp[mask], mask)
		}

		tmp := all ^ mask
		if marked[tmp] {
			best = max(best, tmp*dp[mask])
		}

		for i := 0; i < h; i++ {
			if (mask>>i)&1 == 0 {
				dp[mask|(1<<i)] = max(dp[mask|(1<<i)], dp[mask])
			}
		}
	}
	return int64(best)
}
