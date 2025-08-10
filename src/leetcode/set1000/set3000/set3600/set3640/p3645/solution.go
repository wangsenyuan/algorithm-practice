package p3645

import (
	"cmp"
	"container/heap"
	"slices"
)

func maxTotal(value []int, limit []int) int64 {
	type pair struct {
		first  int
		second int
	}
	n := len(value)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{limit[i], value[i]}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, b.second-a.second)
	})

	var active IntHeap

	var res int
	var i int
	for i < n {
		res += arr[i].second
		heap.Push(&active, arr[i].first)
		i++
		// 目前激活的数量
		sz := len(active)
		for len(active) > 0 && active[0] <= sz {
			heap.Pop(&active)
		}
		for i < n && arr[i].first <= sz {
			i++
		}
	}

	return int64(res)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
