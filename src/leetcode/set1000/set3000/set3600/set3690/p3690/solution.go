package p3690

import (
	"container/heap"
)

type pair struct {
	first  int
	second int
}

func maxTotalValue(nums []int, k int) int64 {
	// 最大值和最小值组成的区间，肯定需要被选中
	// k -= x
	n := len(nums)

	s1 := NewSegTree(n, 0, func(a, b int) int {
		return max(a, b)
	})

	s2 := NewSegTree(n, 1<<60, func(a, b int) int {
		return min(a, b)
	})

	for i, num := range nums {
		s1.Update(i, num)
		s2.Update(i, num)
	}

	seen := make(map[pair]bool)

	var pq PriorityQueue

	add := func(l int, r int) {
		key := pair{l, r}
		if !seen[key] {
			it := new(Item)
			it.id = l*n + r
			it.priority = s1.Get(l, r+1) - s2.Get(l, r+1)
			heap.Push(&pq, it)
			seen[key] = true
		}
	}

	add(0, n-1)

	// 这里k比较小，是个突破口
	var res int64
	for k > 0 {
		it := heap.Pop(&pq).(*Item)
		res += int64(it.priority)
		k--
		l, r := it.id/n, it.id%n
		if l+1 <= r {
			add(l+1, r)
			add(l, r-1)
		}
	}

	return res
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

type SegTree struct {
	arr []int
	sz  int
	iv  int
	f   func(int, int) int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, f}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] = v
	for p > 1 {
		t.arr[p>>1] = t.f(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	res := t.iv
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res = t.f(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = t.f(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
