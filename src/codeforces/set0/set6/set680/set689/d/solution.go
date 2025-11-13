package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve1(a []int, b []int) int {
	n := len(a)

	s1 := NewSegTree(n, -inf, func(a, b int) int {
		return max(a, b)
	})

	s2 := NewSegTree(n, inf, func(a, b int) int {
		return min(a, b)
	})

	for i := range n {
		s1.Update(i, a[i])
		s2.Update(i, b[i])
	}

	check := func(l int) int {

		r1 := sort.Search(n, func(r int) bool {
			return s1.Get(l, r+1) > s2.Get(l, r+1)
		})

		r2 := sort.Search(n, func(r int) bool {
			return s1.Get(l, r+1) >= s2.Get(l, r+1)
		})

		return r1 - r2
	}

	var ans int

	for i := range n {
		ans += check(i)
	}

	return ans
}

const inf = 1 << 60

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = seg.op(v, seg.arr[p])
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func solve(a []int, b []int) int {
	n := len(a)
	d1 := NewDeque(n)
	d2 := NewDeque(n)
	var ans int

	var j int
	for i := range n {
		for !d1.Empty() && a[d1.Back()] <= a[i] {
			d1.PopBack()
		}
		for !d2.Empty() && b[d2.Back()] >= b[i] {
			d2.PopBack()
		}
		d1.PushBack(i)
		d2.PushBack(i)
		for j <= i && a[d1.Front()] > b[d2.Front()] {
			if d1.Front() == j {
				d1.PopFront()
			}
			if d2.Front() == j {
				d2.PopFront()
			}
			j++
		}

		if !d1.Empty() && !d2.Empty() && a[d1.Front()] == b[d2.Front()] {
			ans += min(d1.Front(), d2.Front()) - j + 1
		}
	}

	return ans
}

type Deque struct {
	arr        []int
	head, tail int
}

func NewDeque(n int) *Deque {
	arr := make([]int, n)
	var head, tail int
	return &Deque{arr, head, tail}
}

func (dq *Deque) Empty() bool {
	return dq.head == dq.tail
}

func (dq *Deque) Front() int {
	return dq.arr[dq.tail]
}

func (dq *Deque) Back() int {
	return dq.arr[dq.head-1]
}

func (dq *Deque) PushBack(v int) {
	dq.arr[dq.head] = v
	dq.head++
}

func (dq *Deque) PopBack() int {
	v := dq.arr[dq.head-1]
	dq.head--
	return v
}

func (dq *Deque) PopFront() int {
	v := dq.arr[dq.tail]
	dq.tail++
	return v
}
