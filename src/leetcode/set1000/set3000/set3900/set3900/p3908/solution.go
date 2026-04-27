package p3908

func validDigit(n int, x int) bool {
	found := false
	for n >= 10 {
		if n%10 == x {
			found = true
		}
		n /= 10
	}
	return n != x && found
}
