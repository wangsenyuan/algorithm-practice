package p3946

func maximumSaleItems(items [][]int, budget int) int {
	// n := len(items)

	minPrice := items[0][1]
	maxFactor := items[0][0]
	for _, cur := range items {
		minPrice = min(minPrice, cur[1])
		maxFactor = max(maxFactor, cur[0])
	}

	free := make([]int, maxFactor+1)
	for _, cur := range items {
		free[cur[0]]++
	}

	for i := 1; i <= maxFactor; i++ {
		for j := 2 * i; j <= maxFactor; j += i {
			free[i] += free[j]
		}
	}

	dp := make([]int, budget+1)
	for i := range budget + 1 {
		dp[i] = -inf
	}
	dp[0] = 0

	// 如果购买物品i, 那么额外可以获得 free[factor[i]]个
	// dp[i] 是在每个item只购买一次的情况下，最优解
	for _, cur := range items {
		f, p := cur[0], cur[1]
		for w := budget; w >= p; w-- {
			dp[w] = max(dp[w], dp[w-p]+free[f])
		}
	}
	best := budget / minPrice

	for w := range budget + 1 {
		tmp := dp[w] + (budget-w)/minPrice
		best = max(best, tmp)
	}

	return best
}

const inf = 1 << 60
