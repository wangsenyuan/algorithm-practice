package p3658

func gcdOfOddEvenSums(n int) int {
	s1 := (1 + 2*n - 1) * n / 2
	s2 := (2 + 2*n) * n / 2
	return gcd(s1, s2)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
