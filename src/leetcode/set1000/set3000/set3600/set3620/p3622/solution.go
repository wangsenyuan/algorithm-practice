package p3622

func checkDivisibility(n int) bool {
	var sum int
	prod := 1
	for num := n; num > 0; num /= 10 {
		sum += num % 10
		prod *= num % 10
	}
	return n%(sum+prod) == 0
}
