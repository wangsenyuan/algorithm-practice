package p3867

import "slices"

func gcdSum(nums []int) int64 {
	n := len(nums)
	pref := make([]int, n)
	var mx int
	for i, v := range nums {
		mx = max(mx, v)
		pref[i] = gcd(mx, v)
	}
	slices.Sort(pref)

	var res int
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		res += gcd(pref[l], pref[r])
	}

	return int64(res)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
