package p3754

func sumAndMultiply(n int) int64 {
	var sum int
	var num int
	pw := 1
	for i := n; i > 0; i /= 10 {
		r := i % 10
		sum += r
		if r != 0 {
			num += r * pw
			pw *= 10
		}
	}
	return int64(sum) * int64(num)
}
