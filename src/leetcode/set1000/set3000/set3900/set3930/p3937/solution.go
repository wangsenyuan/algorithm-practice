package p3937

func minOperations(nums []int, k int) int {
	dp := make([][]int, 2)
	for i := range 2 {
		dp[i] = make([]int, k)
	}
	for i, x := range nums {
		for v := range dp[i%2] {
			w := abs(x%k - v)
			dp[i%2][v] += min(w, k-w)
		}
	}

	ans := 1 << 60

	for x := range k {
		for y := range k {
			if x != y {
				ans = min(ans, dp[0][x]+dp[1][y])
			}
		}
	}

	return ans
}

func abs(x int) int {
	return max(x, -x)
}
