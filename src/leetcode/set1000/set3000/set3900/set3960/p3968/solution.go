package p3968

func maxDistance(moves string) int {
	var res int
	var dx, dy int
	for i := range moves {
		if moves[i] == 'U' {
			dx++
		} else if moves[i] == 'D' {
			dx--
		} else if moves[i] == 'L' {
			dy--
		} else if moves[i] == 'R' {
			dy++
		} else {
			res++
		}
	}

	return res + abs(dx) + abs(dy)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
