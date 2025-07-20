package p3623

import (
	"cmp"
	"slices"
)

const mod = 1_000_000_007

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func nC2(n int) int {
	if n%2 == 0 {
		return mul(n/2, n-1)
	}
	return mul(n, (n-1)/2)
}

func countTrapezoids(points [][]int) int {
	// y = b的那条线
	slices.SortFunc(points, func(a []int, b []int) int {
		return cmp.Or(b[1]-a[1], a[0]-b[0])
	})
	// 按照y降序排
	n := len(points)

	var ans int
	var cnt int

	for i := 0; i < n; {
		j := i
		for i < n && points[i][1] == points[j][1] {
			i++
		}
		tmp := nC2(i - j)
		ans = add(ans, mul(tmp, cnt))
		cnt = add(cnt, tmp)
	}

	return ans
}
