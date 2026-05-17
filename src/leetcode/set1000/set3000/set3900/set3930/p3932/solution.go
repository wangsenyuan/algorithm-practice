package p3932

import "sort"

func countKthRoots(l int, r int, k int) int {
	if k == 1 {
		return r - l + 1
	}
	if r == 0 {
		return 1
	}

	check := func(x int, y int) bool {
		if x == 0 {
			return false
		}
		// pow(x, k) > r
		pw := 1
		for range k {
			if pw > y/x || pw*x > y {
				return true
			}
			pw *= x
		}
		return false
	}

	hi := sort.Search(r, func(x int) bool {
		return check(x, r)
	})
	hi--

	if l == 0 {
		return hi + 1
	}
	// pow(hi, k) <= r
	lo := sort.Search(l, func(x int) bool {
		return check(x, l-1)
	})
	if hi < lo {
		return 0
	}

	return hi - lo + 1
}
