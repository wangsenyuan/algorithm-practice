package p3628

func numOfSubsequences(s string) int64 {
	n := len(s)
	pref := make([][]int, n)
	// L, LC, LCT
	for i := 0; i < n; i++ {
		pref[i] = make([]int, 3)
		if i > 0 {
			copy(pref[i], pref[i-1])
		}
		if s[i] == 'L' {
			pref[i][0]++
		} else if s[i] == 'C' {
			pref[i][1] += pref[i][0]
		} else if s[i] == 'T' {
			pref[i][2] += pref[i][1]
		}
	}
	var ans int
	suf := make([]int, 3)
	for i := n - 1; i >= 0; i-- {
		// 如果在这里放置一个T
		ans = max(ans, pref[i][1])
		// 如果这里放置一个C
		ans = max(ans, pref[i][0]*suf[0])
		// 如果在这里放置一个L
		ans = max(ans, suf[1])
		if s[i] == 'T' {
			suf[0]++
		} else if s[i] == 'C' {
			suf[1] += suf[0]
		} else if s[i] == 'L' {
			suf[2] += suf[1]
		}
	}

	ans += pref[n-1][2]
	ans = max(ans, suf[1]+suf[2])

	return int64(ans)
}
