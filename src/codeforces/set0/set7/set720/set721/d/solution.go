package main

import (
	"bufio"
	"cmp"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, b := drive(reader)
	s := fmt.Sprintf("%v", b)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a []int, k int, x int, b []int) {
	var n int
	fmt.Fscan(reader, &n, &k, &x)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b = solve(slices.Clone(a), k, x)
	return
}

type data struct {
	id  int
	val int
}

const inf = 1 << 60

func solve(a []int, k int, x int) []int {
	n := len(a)
	arr := make([]data, n)
	for i, v := range a {
		arr[i] = data{i, v}
	}

	slices.SortFunc(arr, func(x data, y data) int {
		return cmp.Or(x.val-y.val, x.id-y.id)
	})

	var pos int
	for pos < n && arr[pos].val < 0 {
		pos++
	}

	neg := pos
	for pos < n && arr[pos].val == 0 {
		pos++
	}

	zero := pos - neg

	if zero > k {
		// 结果始终是0
		return a
	}

	k -= zero

	if zero > 0 {
		for i := neg; i < pos; i++ {
			arr[i].val = x
		}
		pos = neg
		if neg&1 == 0 {
			arr[neg].val = -x
			pos++
		}
	}

	var pq PriorityQueue
	for i, cur := range arr {
		it := new(Item)
		it.id = i
		it.priority = abs(cur.val)
		heap.Push(&pq, it)
	}

	for k > 0 {
		it := heap.Pop(&pq).(*Item)
		i := it.id
		if pos&1 == 0 {
			// 目前还是正数
			v := it.priority
			c := min(v/x+1, k)
			it.priority = abs(v - c*x)
			k -= c
			if arr[i].val < 0 {
				arr[i].val += c * x
			} else {
				arr[i].val -= c * x
			}
			pos++
		} else {
			it.priority += x
			if arr[i].val > 0 {
				arr[i].val += x
			} else {
				arr[i].val -= x
			}
			k--
		}
		heap.Push(&pq, it)
	}

	for i := range n {
		a[arr[i].id] = arr[i].val
	}

	return a
}

func abs(num int) int {
	return max(num, -num)
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
	return pq[i].priority < pq[j].priority
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
