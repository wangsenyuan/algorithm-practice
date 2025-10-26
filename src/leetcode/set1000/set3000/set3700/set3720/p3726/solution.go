package p3726

func removeZeros(n int64) int64 {
	var res int64
	var base int64 = 1
	for n > 0 {
		x := n % 10
		if x > 0 {
			res += x * base
			base *= 10
		}
		n /= 10
	}

	return res
}
