package p3961

func maxRatings(units [][]int) int64 {
	m := len(units)
	n := len(units[0])
	if n == 1 {
		var res int
		for i := range m {
			res += units[i][0]
		}
		return int64(res)
	}

	var score int
	firstMinVal := 1 << 60
	secondMinVal := 1 << 60
	for i := range m {
		if units[i][0] > units[i][1] {
			units[i][0], units[i][1] = units[i][1], units[i][0]
		}
		for j := 2; j < n; j++ {
			x := units[i][j]
			if x <= units[i][0] {
				x, units[i][0] = units[i][0], x
			}
			if x <= units[i][1] {
				x, units[i][1] = units[i][1], x
			}
		}
		firstMinVal = min(firstMinVal, units[i][0])
		secondMinVal = min(secondMinVal, units[i][1])

		score += units[i][1]
	}

	score -= secondMinVal
	score += firstMinVal

	return int64(score)
}
