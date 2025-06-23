package p3588

import (
	"cmp"
	"slices"
)

func maxArea(coords [][]int) int64 {
	// ˝n := len(coords)

	slices.SortFunc(coords, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	var ans int

	n := len(coords)

	for i := 0; i < n; {
		j := i
		for i < n && coords[i][0] == coords[j][0] {
			i++
		}
		if i-j > 1 {
			// 至少有两个点
			height := coords[i-1][1] - coords[j][1]
			width := max(coords[n-1][0]-coords[j][0], coords[j][0]-coords[0][0])
			ans = max(ans, height*width)
		}
	}

	slices.SortFunc(coords, func(a, b []int) int {
		return cmp.Or(a[1]-b[1], a[0]-b[0])
	})

	for i := 0; i < n; {
		j := i
		for i < n && coords[i][1] == coords[j][1] {
			i++
		}
		if i-j > 1 {
			width := coords[i-1][0] - coords[j][0]
			height := max(coords[n-1][1]-coords[j][1], coords[j][1]-coords[0][1])
			ans = max(ans, height*width)
		}
	}

	if ans == 0 {
		return -1
	}

	return int64(ans)
}
