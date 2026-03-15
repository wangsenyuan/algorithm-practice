package p3870

func countCommas(n int) int {
	var res int

	// 先找到最小的lo
	var cnt int
	lo := 1
	for lo <= n/1000 && lo*1000 <= n {
		cnt++
		lo *= 1000
	}

	num := n
	for num >= 1000 {
		hi := n + 1
		if lo <= hi/1000 {
			hi = lo * 1000
		}
		res += (hi - lo) * cnt
		num = lo / 1000
		lo /= 1000
		cnt--
	}

	return res
}
