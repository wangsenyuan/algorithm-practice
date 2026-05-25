package p1871

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	dp := make([]int, n)
	for i := range n {
		if i > 0 {
			dp[i] += dp[i-1]
		}
		if i == 0 || dp[i] > 0 && s[i] == '0' {
			j := i + minJump
			k := i + maxJump + 1
			if j < n {
				dp[j]++
			}
			if k < n {
				dp[k]--
			}
		}
	}
	return dp[n-1] > 0
}
