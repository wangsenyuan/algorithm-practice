package p3871

func countCommas(n int64) int64 {
	var res int

	// 先找到最小的lo
	var cnt int
	lo := 1
	for lo <= int(n)/1000 && lo*1000 <= int(n) {
		cnt++
		lo *= 1000
	}

	num := int(n)
	for num >= 1000 {
		hi := int(n) + 1
		if lo <= hi/1000 {
			hi = lo * 1000
		}
		res += (hi - lo) * cnt
		num = lo / 1000
		lo /= 1000
		cnt--
	}
	return int64(res)
}
