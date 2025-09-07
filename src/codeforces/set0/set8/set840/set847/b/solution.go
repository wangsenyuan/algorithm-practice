package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		for _, v := range cur {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

func solve(a []int) [][]int {
	arr := slices.Clone(a)
	sort.Ints(arr)

	n := len(arr)

	tr := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})

	seq := make([]int, n)

	for i, v := range a {
		j := sort.SearchInts(arr, v)
		seq[i] = tr.Get(j, n) + 1
		tr.Update(j, seq[i])
	}
	x := tr.Get(0, n)
	res := make([][]int, x+1)

	for i := range n {
		res[seq[i]] = append(res[seq[i]], a[i])
	}

	return res
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
	seg.arr[p] = v
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
