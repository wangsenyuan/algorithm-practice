package p3547

func maxScore(n int, edges [][]int) int64 {
	ans := (n*n*2 + n*5 - 6) * (n - 1) / 6
	if n == len(edges) { // 环
		ans += 2
	}
	return int64(ans)
}
