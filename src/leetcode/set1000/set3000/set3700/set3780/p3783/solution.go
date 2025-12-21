package p3783

func mirrorDistance(n int) int {
	var m int
	for i := n; i > 0; i /= 10 {
		m = m*10 + i%10
	}
	return max(n-m, m-n)
}
