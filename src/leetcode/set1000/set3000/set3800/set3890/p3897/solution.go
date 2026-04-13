package p3897

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func maxValue(nums1 []int, nums0 []int) int {
	var special []pair
	var other []pair
	var mxLen int
	for i := range len(nums1) {
		if nums0[i] == 0 {
			special = append(special, pair{nums1[i], 0})
		} else {
			other = append(other, pair{nums1[i], nums0[i]})
		}
		mxLen = max(mxLen, nums1[i])
	}

	slices.SortFunc(other, func(a, b pair) int {
		// 1的长度不相同时，使用长的那个，相同时，使用0少的那个
		return cmp.Or(b.first-a.first, a.second-b.second)
	})

	arr := append(special, other...)

	p2 := make([]int, mxLen+1)
	p2[0] = 1
	for i := 1; i <= mxLen; i++ {
		p2[i] = add(mul(p2[i-1], 2), 1)
	}

	var res int
	var totLen int
	for i := len(arr) - 1; i >= 0; i-- {
		cur := arr[i]
		totLen += cur.second
		if cur.first > 0 {
			w := p2[cur.first-1]
			tmp := mul(w, pow(2, totLen))
			res = add(res, tmp)
		}
		totLen += cur.first
	}

	return res
}
