package p3579

func minOperations(word1 string, word2 string) int {
	n := len(word1)

	freq := make([][]int, 26)
	for i := range 26 {
		freq[i] = make([]int, 26)
	}

	get := func(s string, t string) int {
		for i := range 26 {
			clear(freq[i])
		}
		var res int
		for i := range len(s) {
			if s[i] == t[i] {
				continue
			}
			x := int(s[i] - 'a')
			y := int(t[i] - 'a')
			if freq[y][x] > 0 {
				freq[y][x]--
			} else {
				res++
				freq[x][y]++
			}

		}
		return res
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for r := 1; r <= n; r++ {
		dp[r] = r
		for l := r - 1; l >= 0; l-- {
			x := get(word1[l:r], word2[l:r])
			dp[r] = min(dp[r], dp[l]+x)
			x = get(reverse(word1[l:r]), word2[l:r])
			dp[r] = min(dp[r], dp[l]+1+x)
		}
	}
	return dp[n]
}

func reverse(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func abs(num int) int {
	return max(num, -num)
}
