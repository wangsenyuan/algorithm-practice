package p3809

func bestTower(towers [][]int, center []int, radius int) []int {
	res := []int{-1, -1, -1}

	check := func(a []int) bool {
		dx := a[0] - center[0]
		dy := a[1] - center[1]
		return abs(dx)+abs(dy) <= radius
	}

	for _, cur := range towers {
		if check(cur) {
			if cur[2] > res[2] || cur[2] == res[2] && (cur[0] < res[0] || cur[0] == res[0] && cur[1] < res[1]) {
				copy(res, cur)
			}
		}
	}

	return res[:2]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}