package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (workers [][]int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	workers = make([][]int, n)
	for i := range n {
		workers[i] = make([]int, 3)
		fmt.Fscan(reader, &workers[i][0], &workers[i][1], &workers[i][2])
	}
	res = solve(n, workers)
	return
}

type worker struct {
	id int
	l  int
	v  int
	r  int
}

func solve(n int, workers [][]int) []int {
	arr := make([]worker, n)
	for i := range n {
		arr[i] = worker{id: i, l: workers[i][0], v: workers[i][1], r: workers[i][2]}
	}
	// v[i] <= v[j]
	slices.SortFunc(arr, func(a, b worker) int {
		return a.v - b.v
	})

	m := arr[n-1].v + 1
	tr := NewTree(m)

	var pq PQ

	var best int

	for i, cur := range arr {
		l, v, r := cur.l, cur.v, cur.r
		heap.Push(&pq, Item{id: i, priority: r})
		tr.Update(l, v, 1)

		for pq.Len() > 0 && pq[0].priority < v {
			it := heap.Pop(&pq).(Item)
			l1, v1 := arr[it.id].l, arr[it.id].v
			tr.Update(l1, v1, -1)
		}
		best = max(best, tr.val[0])
	}

	for pq.Len() > 0 {
		heap.Pop(&pq)
	}

	tr.Reset()

	play := func(l int) []int {
		var res []int
		// l是最大[l,?]且有最大值的区间起点
		for pq.Len() > 0 {
			it := heap.Pop(&pq).(Item)
			l1, v1 := arr[it.id].l, arr[it.id].v
			if l1 <= l && l <= v1 {
				res = append(res, arr[it.id].id+1)
			}
		}
		return res
	}

	for i, cur := range arr {
		l, v, r := cur.l, cur.v, cur.r
		heap.Push(&pq, Item{id: i, priority: r})
		tr.Update(l, v, 1)
		for pq.Len() > 0 && pq[0].priority < v {
			it := heap.Pop(&pq).(Item)
			l1, v1 := arr[it.id].l, arr[it.id].v
			tr.Update(l1, v1, -1)
		}
		tmp := tr.GetBest()
		if tmp.first == best {
			return play(tmp.second)
		}
	}

	return nil
}

type Item struct {
	id       int
	priority int
}

type PQ []Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x any) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Tree struct {
	lazy []int
	val  []int
	sz   int
}

func NewTree(n int) *Tree {
	lazy := make([]int, 4*n)
	val := make([]int, 4*n)
	return &Tree{lazy, val, n}
}

func (t *Tree) apply(i int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(i*2+1, t.lazy[i])
		t.apply(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		t.pull(i)
	}
	f(0, 0, t.sz-1, L, R)
}

type pair struct {
	first  int
	second int
}

func (t *Tree) GetBest() pair {
	var f func(i int, l int, r int) pair
	f = func(i int, l int, r int) pair {
		if l == r {
			return pair{t.val[i], l}
		}
		t.push(i)
		mid := (l + r) >> 1
		if t.val[i*2+2] == t.val[i] {
			return f(i*2+2, mid+1, r)
		}
		return f(i*2+1, l, mid)
	}
	return f(0, 0, t.sz-1)
}

func (t *Tree) Reset() {
	clear(t.lazy)
	clear(t.val)
}
