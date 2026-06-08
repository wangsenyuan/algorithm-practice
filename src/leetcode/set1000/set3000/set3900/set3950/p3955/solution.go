package p3955

func generateValidStrings(n int, k int) []string {

	check := func(mask int) bool {
		var sum int
		var cnt int
		var i int
		for mask > 0 {
			if mask&1 == 1 {
				cnt++
				sum += i
			} else {
				cnt = 0
			}
			if cnt == 2 {
				return false
			}
			i++
			mask >>= 1
		}
		return sum <= k
	}

	format := func(mask int) string {
		buf := make([]byte, n)
		for i := range n {
			if (mask>>i)&1 == 1 {
				buf[i] = '1'
			} else {
				buf[i] = '0'
			}
		}
		return string(buf)
	}

	var res []string
	for mask := range 1 << n {
		if check(mask) {
			res = append(res, format(mask))
		}
	}

	return res
}
