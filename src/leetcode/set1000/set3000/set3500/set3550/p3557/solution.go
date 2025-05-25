package p3557

const inf = 1 << 30

func maxSubstrings(word string) int {
	n := len(word)
	dp := make([]int, n)

	pos := make([]int, 26)
	for i := range 26 {
		pos[i] = -1
	}
	prev := make([]int, n)
	for i := range n {
		x := int(word[i] - 'a')
		prev[i] = pos[x]
		j := i

		for range 3 {
			if j < 0 {
				break
			}
			j = prev[j]

			if j >= 0 && i-j+1 >= 4 {
				if j > 0 {
					dp[i] = dp[j-1] + 1
				} else {
					dp[i] = 1
				}
				break
			}
		}
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
		pos[x] = i
	}

	return dp[n-1]
}
