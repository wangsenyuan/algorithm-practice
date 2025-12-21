package p3782

func lastInteger(n int64) int64 {
	var a0 = 1
	var d = 1
	t := int(n)
	var flag int
	for t > 1 {
		if flag == 0 {
			// a0 不变
			d *= 2
			t = (t + 1) / 2
		} else {
			a0 = a0 + (t-1)*d
			d *= -1
			d *= 2
			t = (t + 1) / 2
			a0 = a0 + (t-1)*d
			d *= -1
		}

		flag ^= 1
	}
	return int64(a0)
}
