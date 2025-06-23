package p3592

func findCoins(numWays []int) []int {
	// 那些1的比较特殊
	var res []int
	// n := len(numWays)
	n := len(numWays)

	dp := make([]int, n+1)
	count := func(x int) int {
		clear(dp)
		dp[0] = 1
		for _, v := range res {
			for u := 0; u+v <= x; u++ {
				dp[u+v] += dp[u]
			}
		}
		return dp[x]
	}

	for i, v := range numWays {
		i++
		tmp := count(i)
		if tmp == v-1 {
			res = append(res, i)
			continue
		}
		if tmp == v {
			continue
		}
		return nil
	}

	return res
}
