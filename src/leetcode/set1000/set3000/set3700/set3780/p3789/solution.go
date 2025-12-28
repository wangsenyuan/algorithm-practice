package p3789


func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	var res int64
	if cost1 + cost2 > costBoth {
		w := min(need1, need2)
		res += int64(w * costBoth)
		need1 -= w
		need2 -= w
	}

	res += int64(min(cost1, costBoth)) * int64(need1)
	res += int64(min(cost2, costBoth)) * int64(need2)

	return res
}