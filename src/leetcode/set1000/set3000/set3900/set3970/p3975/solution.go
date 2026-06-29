package p3975

import (
	"cmp"
	"slices"
)

func filterOccupiedIntervals(occupiedIntervals [][]int, freeStart int, freeEnd int) [][]int {
	slices.SortFunc(occupiedIntervals, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	var res [][]int

	n := len(occupiedIntervals)
	for i := 0; i < n; {
		l, r := occupiedIntervals[i][0], occupiedIntervals[i][1]
		for i < n && occupiedIntervals[i][0] <= r+1 {
			r = max(r, occupiedIntervals[i][1])
			i++
		}
		if freeStart <= l && r <= freeEnd {
			continue
		}
		if r < freeStart || freeEnd < l {
			res = append(res, []int{l, r})
		} else {
			if l < freeStart {
				res = append(res, []int{l, freeStart - 1})
			}
			if freeEnd < r {
				res = append(res, []int{freeEnd + 1, r})
			}
		}
	}

	return res

}
