package p3530

import "math/bits"

const inf = 1 << 60

func maxProfit(n int, edges [][]int, score []int) int {
	dep := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		dep[v] |= 1 << u
	}
	N := 1 << n
	dp := make([]int, N)
	for i := range dp {
		dp[i] = -inf
	}
	dp[0] = 0

	for state := 0; state < N; state++ {
		if dp[state] < 0 {
			continue
		}
		sz := bits.OnesCount(uint(state))
		for i := range n {
			if (state>>i)&1 == 0 && dep[i]&state == dep[i] {
				dp[state|(1<<i)] = max(dp[state|(1<<i)], dp[state]+(sz+1)*score[i])
			}
		}
	}

	return dp[N-1]
}
