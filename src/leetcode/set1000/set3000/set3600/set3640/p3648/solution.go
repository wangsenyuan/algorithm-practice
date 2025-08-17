package p3648

func minSensors(n int, m int, k int) int {
	k = 2*k + 1
	x := (n + k - 1) / k
	y := (m + k - 1) / k
	return x * y
}
