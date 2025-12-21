package p3784

func minCost(s string, cost []int) int64 {
	var sum int
	n := len(s)

	cnt := make([]int, 26)

	for i := range n {
		sum += cost[i]
		x := int(s[i] - 'a')
		cnt[x] += cost[i]
	}

	best := sum
	for i := range 26 {
		best = min(best, sum-cnt[i])
	}
	return int64(best)
}
