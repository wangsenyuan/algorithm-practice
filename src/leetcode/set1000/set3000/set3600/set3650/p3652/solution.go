package p3652

import "slices"

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	// 选择一个k段的区间
	// 在前半段, 在未改变前，未改变前的，利润 = x， 改变后的利润为y
	// 在后半段... a, b 那么就是选择 b - a + y - x 最大的一段
	profit := make([][]int, n+1)
	profit[0] = make([]int, 3)
	for i := range n {
		profit[i+1] = slices.Clone(profit[i])
		profit[i+1][strategy[i]+1] += prices[i]
	}

	ans := -profit[n][0] + profit[n][2]

	var best int

	h := k / 2
	for i := 0; i+k <= n; i++ {
		// 前半个区间花费的费用 b - a
		a := profit[i+h][0] - profit[i][0]
		// 原来卖出的，现在要变成持有
		b := profit[i+h][2] - profit[i][2]

		// 原来买入的，要变成卖出
		c := profit[i+k][0] - profit[i+h][0]
		// 原来持有的，要变成卖出
		d := profit[i+k][1] - profit[i+h][1]

		tmp := d + 2*c + a - b

		best = max(best, tmp)
	}

	return int64(ans + best)
}
