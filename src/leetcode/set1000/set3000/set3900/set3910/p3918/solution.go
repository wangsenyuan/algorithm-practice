package p3918

func sumOfPrimesInRange(n int) int {
	var r int
	for i := n; i > 0; i /= 10 {
		r = r*10 + i%10
	}
	a := min(r, n)
	b := max(r, n)
	set := make([]bool, b+1)
	var primes []int
	for i := 2; i <= b; i++ {
		if !set[i] {
			set[i] = true
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > b {
				break
			}
			set[i*j] = true
			if i%j == 0 {
				break
			}
		}
	}
	var sum int
	for _, p := range primes {
		if p >= a {
			sum += p
		}
	}
	return sum
}
