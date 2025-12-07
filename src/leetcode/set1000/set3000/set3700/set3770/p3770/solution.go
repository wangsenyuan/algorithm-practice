package p3770

func largestPrime(n int) int {
	var primes []int
	marked := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !marked[i] {
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j > n {
				break
			}
			marked[i*j] = true
			if i%j == 0 {
				break
			}
		}
	}

	var ans int

	var sum int

	for _, v := range primes {
		sum += v
		if sum > n {
			break
		}
		if !marked[sum] {
			ans = sum
		}
	}

	return ans
}
