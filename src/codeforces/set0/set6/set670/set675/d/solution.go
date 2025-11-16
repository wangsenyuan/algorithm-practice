package main

import (
	"bufio"
	. "fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := Sprintf("%v", res)
	Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) []int {
	n := len(a)
	children := make([]int, n)
	// 如果要用avl tree吗？
	// 实在是太重了
	t1 := NewSegTree(n, -1, func(a int, b int) int {
		return max(a, b)
	})
	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	ans := make([]int, n)

	pos := make([]int, n)
	for i := range n {
		pos[arr[i].second] = i
	}

	t1.Update(pos[0], pos[0])
	t2.Update(pos[0], pos[0])

	for i := 1; i < n; i++ {
		// 要找到前面比a[i]小的数,
		// 已经前面比a[i]大的数的位置
		l := t1.Get(0, pos[i])
		r := t2.Get(pos[i], n)
		// l和r肯定会找到一个
		if l >= 0 && children[arr[l].second]&2 == 0 {
			p := arr[l].second
			ans[i] = a[p]
			children[p] |= 2
		} else {
			p := arr[r].second
			ans[i] = a[p]
			children[p] |= 1
		}
		t1.Update(pos[i], pos[i])
		t2.Update(pos[i], pos[i])
	}

	return ans[1:]
}

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
