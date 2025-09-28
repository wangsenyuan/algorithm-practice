package p3696

import "slices"

func decimalRepresentation(n int) []int {
	var res []int
	base := 1
	for n > 0 {
		x := n % 10
		if x > 0 {
			res = append(res, x*base)
		}
		n /= 10
		base *= 10
	}
	slices.Reverse(res)
	return res
}
