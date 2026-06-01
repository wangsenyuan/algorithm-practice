package p3947

import (
	"slices"
)

func maximumSaleItems(items [][]int, budget int) int {
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

	slices.SortFunc(items, func(a []int, b []int) int {
		return a[1] - b[1]
	})

	best := budget / minPrice

	var sum int
	var cnt int
	for _, cur := range items {
		w := min(free[cur[0]]-1, (budget-sum)/cur[1])
		sum += w * cur[1]
		cnt += w

		tmp := cnt*2 + (budget-sum)/minPrice

		best = max(best, tmp)
	}

	return best
}
