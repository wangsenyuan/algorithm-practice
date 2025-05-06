package p3538

import "slices"

const inf = 1 << 60

func minTravelTime(l int, n int, k int, position []int, time []int) int {
	// dp[i][j] 表示在i处进行了j次合并的最优解
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, k+1)
		for j := range k + 1 {
			dp[i][j] = make([]int, k+1)
			for u := range k + 1 {
				dp[i][j][u] = inf
			}
		}
	}

	dp[0][0][0] = 0

	pref := make([]int, n+1)
	for i := range n {
		pref[i+1] = pref[i] + time[i]
	}

	for i := 1; i < n; i++ {

		for j := 1; j <= i; j++ {
			// 将j到i处合并到i, 合并了x次
			x := i - j
			// j1处的状态确定了
			j1 := j - 1
			for u := 0; u+x <= k; u++ {
				for v := 0; v <= u; v++ {
					if dp[j1][u][v] != inf {
						tmp := pref[j1+1] - pref[j1-v]
						dp[i][u+x][x] = min(dp[i][u+x][x], dp[j1][u][v]+tmp*(position[i]-position[j1]))
					}
				}
			}
		}
	}

	return slices.Min(dp[n-1][k])
}
