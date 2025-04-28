package p3531

import (
	"cmp"
	"slices"
)

func countCoveredBuildings(m int, buildings [][]int) int {
	slices.SortFunc(buildings, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})
	n := len(buildings)

	prev := make(map[int]int)

	cnt := make([]int, n)

	for i := 0; i < n; {
		j := i
		for i < n && buildings[i][0] == buildings[j][0] {
			// 同一条垂线
			if i > j {
				cnt[i-1]++
				cnt[i]++
			}
			if k, ok := prev[buildings[i][1]]; ok {
				cnt[k]++
				cnt[i]++
			}
			prev[buildings[i][1]] = i
			i++
		}
	}

	var res int

	for i := range n {
		if cnt[i] == 4 {
			res++
		}
	}

	return res
}
