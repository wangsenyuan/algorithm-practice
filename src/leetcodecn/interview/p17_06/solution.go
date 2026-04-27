package p1706

import "slices"

func numberOf2sInRange(n int) int {
	var digits []int
	for i := n; i > 0; i /= 10 {
		digits = append(digits, i%10)
	}
	var bases []int
	bases = append(bases, 1)
	for range digits {
		bases = append(bases, 10*bases[len(bases)-1])
	}

	slices.Reverse(digits)
	// slices.Reverse(bases)
	// slices.Reverse(lo)

	m := len(digits)

	var res int
	var hi int
	for i, v := range digits {
		// 目前为止高位部分有多少呢？
		lo := n - (hi*10+v)*bases[m-1-i]
		if i > 0 {
			// 把所有小于x高位部分的计算进去
			res += hi * bases[m-1-i]
		}
		// 然后考虑和高位部分相等的值
		if 2 < v {
			res += bases[m-1-i]
		} else if 2 == v {
			res += lo + 1
		}

		hi = hi*10 + v
	}

	return res
}
