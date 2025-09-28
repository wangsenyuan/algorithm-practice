package p3694

type pair struct {
	first  int
	second int
}

func distinctPoints(s string, k int) int {
	n := len(s)
	dp := make([]pair, n+1)
	var x, y int
	for i := range n {
		if s[i] == 'U' {
			y++
		} else if s[i] == 'D' {
			y--
		} else if s[i] == 'L' {
			x--
		} else {
			x++
		}
		dp[i+1] = pair{x, y}
	}
	freq := make(map[pair]int)

	add := func(l int, r int) {
		first := dp[l]
		second := dp[r]
		x, y := first.first, first.second
		dx := dp[n].first - second.first
		dy := dp[n].second - second.second
		freq[pair{x + dx, y + dy}]++
	}

	for i := k; i <= n; i++ {
		add(i-k, i)
	}

	return len(freq)
}
