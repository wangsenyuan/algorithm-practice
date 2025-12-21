package p3781

import (
	"container/heap"
)

func maximumScore(nums []int, s string) int64 {
	// 也就是说1，只能往前移动
	n := len(nums)
	var res int

	var pq IntHeap

	for i := 0; i < n; i++ {
		heap.Push(&pq, nums[i])
		if s[i] == '1' {
			// 遇到1，就把目前最大的数拿出来用
			res += heap.Pop(&pq).(int)
		}
	}
	return int64(res)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
