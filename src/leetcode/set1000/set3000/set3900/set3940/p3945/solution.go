package p3945

func digitFrequencyScore(n int) int {
	var res int
	for i := n; i > 0; i /= 10 {
		res += i % 10
	}
	return res
}
