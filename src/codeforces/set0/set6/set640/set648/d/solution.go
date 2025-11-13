package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	dogs := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &dogs[i])
	}
	bowls := make([][]int, m)
	for i := range m {
		bowls[i] = make([]int, 2)
		fmt.Fscan(reader, &bowls[i][0], &bowls[i][1])
	}
	return solve(dogs, bowls)
}

func solve(dogs []int, bowls [][]int) int {
	slices.Sort(dogs)

	slices.SortFunc(bowls, func(a []int, b []int) int {
		return (a[0] - a[1]) - (b[0] - b[1])
	})

	var ans int
	var pq IntHeap

	n := len(dogs)
	m := len(bowls)

	for i, j := 0, 0; i < n; i++ {
		for len(pq) > 0 && pq[0] < dogs[i] {
			heap.Pop(&pq)
		}

		for j < m && bowls[j][0]-bowls[j][1] <= dogs[i] {
			r := bowls[j][0] + bowls[j][1]
			if r >= dogs[i] {
				heap.Push(&pq, r)
			}
			j++
		}
		if len(pq) > 0 {
			ans++
			heap.Pop(&pq)
		}
	}

	return ans
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

func solve1(dogs []int, bowls [][]int) int {
	sort.Ints(dogs)

	n := len(dogs)

	var xs []int

	xs = append(xs, dogs...)
	for _, bowl := range bowls {
		l, r := bowl[0]-bowl[1], bowl[0]+bowl[1]
		xs = append(xs, l, r)
	}

	slices.Sort(xs)
	xs = slices.Compact(xs)

	m := len(xs)

	open := make([][]int, m)

	for i, cur := range bowls {
		l := cur[0] - cur[1]
		l = sort.SearchInts(xs, l)
		open[l] = append(open[l], i)
	}

	// 貌似用一个heap也可以
	t := NewTree(m)

	var ans int

	var p int
	for i := range m {
		for _, j := range open[i] {
			r := bowls[j][1] + bowls[j][0]
			r = sort.SearchInts(xs, r)
			t.Update(r, 1)
		}
		for p < n && dogs[p] == xs[i] {
			j := t.Find()
			if j != -1 {
				ans++
				t.Update(j, -1)
			}
			p++
		}
		// 还在位置i的，都是过期的
		c := t.Get(i)
		if c > 0 {
			t.Update(i, -c)
		}
	}

	return ans
}

type Tree struct {
	cnt []int
}

func NewTree(n int) *Tree {
	cnt := make([]int, 4*n)
	return &Tree{cnt: cnt}
}

func (t *Tree) Update(p int, v int) {
	n := len(t.cnt) / 4
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.cnt[i] += v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}
	f(0, 0, n-1)
}

func (t *Tree) Get(p int) int {
	n := len(t.cnt) / 4
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return t.cnt[i]
		}
		mid := (l + r) / 2
		if p <= mid {
			return f(i*2+1, l, mid)
		} else {
			return f(i*2+2, mid+1, r)
		}
	}
	return f(0, 0, n-1)
}

func (t *Tree) Find() int {
	if t.cnt[0] == 0 {
		return -1
	}

	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if t.cnt[i*2+1] > 0 {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	n := len(t.cnt) / 4
	return f(0, 0, n-1)
}
