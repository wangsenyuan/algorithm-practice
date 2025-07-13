package p3614

func processStr(s string, K int64) byte {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		x := s[i-1]
		if x == '*' {
			dp[i] = max(dp[i-1]-1, 0)
		} else if x == '#' {
			dp[i] = dp[i-1] * 2
		} else if x == '%' {
			// reverse
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	k := int(K + 1)
	if k > dp[n] {
		return '.'
	}

	for i := n; i > 0; i-- {
		x := s[i-1]
		if x >= 'a' && x <= 'z' {
			if k == dp[i] {
				return x
			}
			// k < dp[i]
			continue
		}
		if x == '%' {
			// reverse
			k = dp[i] - k + 1
		} else if x == '#' {
			// double
			if k > dp[i-1] {
				k -= dp[i-1]
			}
		}
		// else remove one
	}

	return '.'
}
