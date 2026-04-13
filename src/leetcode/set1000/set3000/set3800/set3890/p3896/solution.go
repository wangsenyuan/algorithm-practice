package p3896

import (
	"slices"
	"sort"
)

func minOperations(nums []int) int {
	x := slices.Max(nums)
	x *= 3
	lpf := make([]int, x+1)
	var primes []int
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > x {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	var res int
	for i, v := range nums {
		if i&1 == 0 {
			if lpf[v] != v {
				w := sort.SearchInts(primes, v+1)
				res += primes[w] - v
			}
		} else {
			for d := 0; ; d++ {
				if v+d <= x && lpf[v+d] != v+d {
					res += d
					break
				}
			}
		}
	}

	return res
}
