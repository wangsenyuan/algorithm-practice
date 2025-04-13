package p3516

func findClosest(x int, y int, z int) int {
	a := abs(z - x)
	b := abs(z - y)
	if a < b {
		return 1
	}
	if a > b {
		return 2
	}
	return 0
}

func abs(num int) int {
	return max(num, -num)
}
