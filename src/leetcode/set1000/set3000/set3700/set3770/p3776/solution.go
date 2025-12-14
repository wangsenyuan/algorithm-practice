package p3776

func minMoves(balance []int) int64 {
	n := len(balance)
	var pos int
	var sum int
	for i := range n {
		sum += balance[i]
		if balance[i] < 0 {
			pos = i
		}
	}
	if sum < 0 {
		return -1
	}

	var res int

	w := -balance[pos]

	for d := 1; w > 0; d++ {
		v := balance[(pos-d+n)%n] + balance[(pos+d)%n]
		x := min(v, w)
		res += d * x
		w -= x
	}

	return int64(res)
}
