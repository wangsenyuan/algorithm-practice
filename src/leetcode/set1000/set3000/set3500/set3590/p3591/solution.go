package p3591

import "slices"

func checkPrimeFrequency(nums []int) bool {
	x := slices.Max(nums)
	freq := make([]int, x+1)
	for _, v := range nums {
		freq[v]++
	}

	for _, v := range freq {
		if isPrime(v) {
			return true
		}
	}
	return false
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
