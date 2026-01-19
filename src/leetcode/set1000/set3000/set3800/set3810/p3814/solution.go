package p3814

import (
	"slices"
	"sort"
)

type item struct {
	cost     int
	capacity int
}

func maxCapacity(costs []int, capacity []int, budget int) int {
	n := len(costs)
	arr := make([]item, n)
	for i := range n {
		arr[i] = item{costs[i], capacity[i]}
	}
	slices.SortFunc(arr, func(a item, b item) int {
		return a.cost - b.cost
	})

	var res int

	best := make([]int, n)
	for i := range n {
		best[i] = arr[i].capacity
		if i > 0 {
			best[i] = max(best[i], best[i-1])
		}
	}

	budget--

	for r := n - 1; r >= 0; r-- {
		if arr[r].cost <= budget {
			res = max(res, arr[r].capacity)
			l := sort.Search(r, func(i int) bool {
				return arr[i].cost+arr[r].cost > budget
			})
			l--
			if l >= 0 && l < r {
				res = max(res, best[l]+arr[r].capacity)
			}
		}
	}

	return res
}
