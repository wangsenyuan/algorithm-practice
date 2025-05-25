package p3563

func lexicographicallySmallestString(s string) string {
	n := len(s)

	fp := make([][]bool, n)
	for i := range n {
		fp[i] = make([]bool, n)
	}
	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			if check(s[l], s[r]) {
				if l+1 == r || fp[l+1][r-1] {
					fp[l][r] = true
				}
			}
			for i := l + 1; i+1 < r && !fp[l][r]; i++ {
				if fp[l][i] && fp[i+1][r] {
					fp[l][r] = true
					break
				}
			}
		}
	}

	dp := make([]string, n+1)
	dp[n] = ""
	for i := n - 1; i >= 0; i-- {
		dp[i] = s[i:i+1] + dp[i+1]
		for j := i + 1; j < n; j += 2 {
			if fp[i][j] {
				dp[i] = min(dp[i], dp[j+1])
			}
		}
	}
	return dp[0]
}

func check(a, b byte) bool {
	x := int(a - 'a')
	y := int(b - 'a')
	x, y = min(x, y), max(x, y)
	return x+1 == y || x+25 == y
}
