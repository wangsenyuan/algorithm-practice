package p3747

import "slices"

func countDistinct(n int64) int64 {
	var digits []int
	for i := n; i > 0; i /= 10 {
		digits = append(digits, int(i%10))
	}

	slices.Reverse(digits)

	m := len(digits)
	p9 := make([]int, m+1)
	p9[0] = 1
	var res int
	for i := 1; i <= m; i++ {
		p9[i] = p9[i-1] * 9
		if i < m {
			res += p9[i]
		}
	}
	bad := false
	// 然后计算m长度的
	for i := 0; i < m; i++ {
		x := digits[i]
		if x == 0 {
			bad = true
			break
		}
		res += p9[m-1-i] * (x - 1)
	}

	if !bad {
		res++
	}

	return int64(res)
}
