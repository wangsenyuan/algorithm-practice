package p3951

import (
	"cmp"
	"slices"
)

func minEnergy(n int, brightness int, intervals [][]int) int64 {
	k := min((brightness+2)/3, n)

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	var res int
	for i := 0; i < len(intervals); {
		l, r := intervals[i][0], intervals[i][1]
		for i < len(intervals) && intervals[i][0] <= r+1 {
			r = max(r, intervals[i][1])
			i++
		}

		res += (r - l + 1) * k
	}

	return int64(res)
}
