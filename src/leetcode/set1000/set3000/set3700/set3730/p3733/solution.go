package p3733

import "sort"

func minimumTime(d []int, r []int) int64 {
	w := lcm(r[0], r[1])
	check := func(n int) bool {
		if n-n/w < d[0]+d[1] {
			return false
		}
		if n-n/r[0] < d[0] {
			return false
		}
		if n-n/r[1] < d[1] {
			return false
		}
		return true
	}

	res := sort.Search(1e18, check)
	return int64(res)
}

func lcm(a, b int) int {
	c := gcd(a, b)
	return a / c * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
