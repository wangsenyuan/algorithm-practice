package main

import "container/heap"

const inf = 1 << 60

func minimumPairRemoval(nums []int) int {
	// 2 1 3
	// 问题麻烦的地方在于怎么判断，数组已经是排好序的
	// a[i] <= a[i] + a[i+1]后
	// cnt += a[i] > a[i+1]
	// cnt = 0 的时候，完成
	n := len(nums)
	if n == 1 {
		return 0
	}
	next := make([]int, n)
	prev := make([]int, n)
	for i := range n {
		next[i] = i + 1
		prev[i] = i - 1
	}

	items := make([]*Item, n-1)
	pq := make(PriorityQueue, n-1)
	var cnt int
	for i := 0; i < n-1; i++ {
		it := new(Item)
		it.id = i
		it.priority = nums[i] + nums[i+1]
		if nums[i] > nums[i+1] {
			cnt++
		}
		it.index = i
		items[i] = it
		pq[i] = it
	}

	arr := make([]int, n)
	copy(arr, nums)

	heap.Init(&pq)

	rem := func(i int) {
		if i < 0 || i == n {
			return
		}
		if next[i] < n && arr[i] > arr[next[i]] {
			// 原来它是计算在内的
			cnt--
		}
	}

	add := func(i int) {
		if i < 0 || i == n {
			return
		}
		if next[i] < n && arr[i] > arr[next[i]] {
			cnt++
		}
	}

	update := func(it *Item) {
		i := it.id
		l := prev[i]
		r := next[i]
		// 只会影响到这三个位置
		rem(r)
		rem(i)
		rem(l)
		if r < n {
			// 先把r给remove掉
			if r < n-1 && items[r].index >= 0 {
				pq.remove(items[r])
			}
			arr[i] += arr[r]
			// 为了debug
			arr[r] = 0
			if next[r] < n {
				next[i] = next[r]
				prev[next[r]] = i
			} else {
				next[i] = n
			}
			if next[i] < n {
				it.priority = arr[i] + arr[next[i]]
				heap.Push(&pq, it)
			}
		}
		// 同时要更新l
		if l >= 0 {
			pq.update(items[l], arr[l]+arr[i])
		}
		add(l)
		add(i)
	}

	var res int
	for cnt > 0 {
		res++
		it := heap.Pop(&pq).(*Item)
		update(it)
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
	return pq[i].priority < pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id < pq[j].id
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

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}

func (pq *PriorityQueue) remove(it *Item) {
	pq.update(it, -inf)
	heap.Pop(pq)
}
