package p3532

import (
	"cmp"
	"slices"
)

type pair struct {
	first  int
	second int
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{nums[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	var m int
	pos := make([]int, n)
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			pos[arr[i].second] = m
			i++
		}
		arr[m] = arr[j]
		m++
	}

	arr = arr[:m]

	next := make([]int, m+1)
	next[m] = m
	for i, r := m-1, m-1; i >= 0; i-- {
		next[i] = i
		for r > i && arr[r].first-arr[i].first > maxDiff {
			r--
		}
		next[i] = next[r]
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		if pos[l] > pos[r] {
			l, r = r, l
		}
		ans[i] = next[pos[l]] >= pos[r]
	}

	return ans
}
