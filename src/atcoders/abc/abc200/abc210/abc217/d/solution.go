package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var L, Q int
	fmt.Fscan(reader, &L, &Q)
	queries := make([][]int, Q)
	for i := range Q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(L, queries)
}

func solve(L int, queries [][]int) []int {
	// L is too large, can't use it directly
	var markers []int
	for _, cur := range queries {
		x := cur[1]
		markers = append(markers, x)
	}
	markers = append(markers, 0, L)
	slices.Sort(markers)
	markers = slices.Compact(markers)
	n := len(markers)
	t1 := NewSegTree(n, -inf, func(a int, b int) int {
		return max(a, b)
	})
	t2 := NewSegTree(n, inf, func(a int, b int) int {
		return min(a, b)
	})

	t1.Update(0, 0)
	t2.Update(n-1, L)

	var ans []int
	for _, cur := range queries {
		x := cur[1]
		i := sort.SearchInts(markers, x)

		if cur[0] == 1 {
			t1.Update(i, x)
			t2.Update(i, x)
		} else {
			// 找到i后面最小的r, 可能找到x
			r := t2.Get(i, n)
			l := t1.Get(0, i)
			ans = append(ans, r-l)
		}
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
