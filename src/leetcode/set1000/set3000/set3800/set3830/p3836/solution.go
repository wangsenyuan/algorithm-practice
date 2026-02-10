package p3836

const inf = 1 << 60

func maxScore(nums1 []int, nums2 []int, k int) int64 {

	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n+1)
	ndp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
		ndp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = -inf
			ndp[i][j] = -inf
		}
	}

	for i := range n {
		clear(dp[i])
	}

	for x := range k {
		for i := x; i < n; i++ {
			for j := x; j < m; j++ {
				ndp[i+1][j+1] = max(dp[i][j]+nums1[i]*nums2[j], ndp[i][j+1], ndp[i+1][j])
			}
		}
		for i := range n + 1 {
			for j := range m + 1 {
				dp[i][j] = ndp[i][j]
				ndp[i][j] = -inf
			}
		}
	}

	return int64(dp[n][m])
}

func maxScore1(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	m := len(nums2)

	dp := make([][][]int, n)
	marked := make([][][]bool, n)
	for i := range n {
		dp[i] = make([][]int, m)
		marked[i] = make([][]bool, m)
		for j := range m {
			marked[i][j] = make([]bool, k+1)
			dp[i][j] = make([]int, k+1)
			for l := range k + 1 {
				dp[i][j][l] = -inf
			}
		}
	}

	var f func(i int, j int, x int) int

	f = func(i int, j int, x int) (res int) {
		if x == 0 {
			return 0
		}

		if i < 0 || j < 0 || min(i, j)+1 < x {
			return -inf
		}

		if marked[i][j][x] {
			return dp[i][j][x]
		}
		marked[i][j][x] = true
		defer func() {
			dp[i][j][x] = res
		}()

		res = max(f(i-1, j, x), f(i, j-1, x))

		res = max(res, f(i-1, j-1, x-1)+nums1[i]*nums2[j])

		return
	}

	res := f(n-1, m-1, k)

	return int64(res)
}
