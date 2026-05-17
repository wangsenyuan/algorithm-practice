package p3931

func isAdjacentDiffAtMostTwo(s string) bool {
	for i := 0; i + 1 < len(s); i++ {
		x := int(s[i] - 'a')
		y := int(s[i+1] - 'a')
		if abs(x - y) > 2 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	return max(x, -x)
}
