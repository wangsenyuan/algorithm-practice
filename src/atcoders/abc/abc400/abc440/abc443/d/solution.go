package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	R := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &R[i])
	}
	return solve(n, R)
}

type pair struct {
	first  int
	second int
}

func solve(n int, R []int) int {
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{n - R[i], i}
	}

	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(y.first-x.first, x.second-y.second)
	})

	// 需要知道两边比自己高的位置

	t1 := NewSegTree(n, -1, func(a int, b int) int {
		return max(a, b)
	})
	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	var res int

	col := make([]int, n)

	for _, cur := range arr {
		i := cur.second

		l := t1.Get(0, i)
		r := t2.Get(i, n)
		x := cur.first
		if l >= 0 {
			x = max(x, col[l]-(i-l))
		}
		if r < n {
			x = max(x, col[r]-(r-i))
		}

		res += x - cur.first

		col[i] = x
		t1.Update(i, i)
		t2.Update(i, i)
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
