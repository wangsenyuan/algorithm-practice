package p3686

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func countStableSubsequences(nums []int) int {
	// dp[0/1][cnt] = 以0/1结尾，且连续有cnt+1个数
	var dp [2][2]int
	for _, v := range nums {
		v = v & 1
		var ndp [2][2]int
		ndp[v][1] = dp[v][0]
		ndp[v][0] = add(1, add(dp[v^1][0], dp[v^1][1]))

		for i := range 2 {
			for j := range 2 {
				dp[i][j] = add(dp[i][j], ndp[i][j])
			}
		}
	}
	var res int
	for i := range 2 {
		for j := range 2 {
			res = add(res, dp[i][j])
		}
	}
	return res
}
