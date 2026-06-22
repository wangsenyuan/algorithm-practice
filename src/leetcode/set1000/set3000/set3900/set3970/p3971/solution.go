package p3971

import (
	"slices"
	"sort"
)

const mod = 1e9 + 7

func add(a, b int) int {
	a %= mod
	b %= mod
	return (a + b) % mod
}

func mul(a, b int) int {
	a %= mod
	b %= mod
	return (a * b) % mod
}

func maxTotalValue(value []int, decay []int, m int) int {
	check := func(low int) bool {
		low++
		var cnt int
		for i, v := range value {
			if v >= low {
				d := decay[i]
				// v - d * (k - 1) >= low
				k := (v-low)/d + 1
				cnt += k
				if cnt > m {
					return false
				}
			}
		}
		return true
	}
	lo := sort.Search(slices.Max(value), check)

	var res int
	for i, v := range value {
		if v > lo {
			d := decay[i]
			k := (v-lo-1)/d + 1
			m -= k
			// v - d * (k - 1)
			x := v - d*(k-1)
			// (x + v) * k / 2
			x += v
			if k%2 == 0 {
				k /= 2
			} else {
				x /= 2
			}

			res = add(res, mul(x, k))
		}
	}

	res = add(res, mul(lo, m))

	return res
}
