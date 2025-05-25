package p3560

func minCuttingCost(n int, m int, k int) int64 {
	res := max(0, min(n, n-k)*k) + max(0, min(m, m-k)*k)
	return int64(res)
}
