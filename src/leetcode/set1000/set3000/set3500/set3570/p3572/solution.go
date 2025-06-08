package p3572

import (
	"cmp"
	"slices"
	"sort"
)

type pair struct {
	first  int
	second int
}

func maxSumDistinctTriplet(x []int, y []int) int {
	n := len(x)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{x[i], y[i]}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	var ys []int

	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			i++
		}
		ys = append(ys, arr[i-1].second)
	}
	sort.Ints(ys)
	if len(ys) < 3 {
		return -1
	}
	m := len(ys)
	return ys[m-2] + ys[m-1] + ys[m-3]
}
