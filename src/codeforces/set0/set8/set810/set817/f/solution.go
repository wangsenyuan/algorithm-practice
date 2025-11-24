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
	res := drive(reader)
	for _, w := range res {
		fmt.Fprintln(writer, w)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	queries := make([][]int, n)
	for i := range n {
		var op, l, r int
		fmt.Fscan(reader, &op, &l, &r)
		queries[i] = []int{op, l, r}
	}
	return solve(queries)
}

func solve(queries [][]int) []int {

	var arr []int
	arr = append(arr, 1)

	for _, cur := range queries {
		l, r := cur[1], cur[2]
		if l > 1 {
			arr = append(arr, l-1)
		}
		arr = append(arr, l, l+1, r, r+1)
		if r > 1 {
			arr = append(arr, r-1)
		}
	}

	sort.Ints(arr)
	arr = append(arr, arr[len(arr)-1]+1)
	arr = slices.Compact(arr)

	n := len(arr)

	tr := NewTree(n)

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[1], cur[2]
		u := sort.SearchInts(arr, l)
		v := sort.SearchInts(arr, r)
		switch cur[0] {
		case 1:
			tr.Update(u, v, 1)
		case 2:
			tr.Update(u, v, -1)
		default:
			tr.Update(u, v, 2)
		}
		w := tr.FindFirstZero()
		ans[i] = arr[w]
	}

	return ans
}

type Tree struct {
	cnt  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	cnt := make([]int, n*4)
	lazy := make([]int, n*4)
	return &Tree{
		cnt:  cnt,
		lazy: lazy,
		sz:   n,
	}
}

func (t *Tree) apply(i int, l int, r int, op int) {
	switch op {
	case 1:
		t.cnt[i] = r - l + 1
	case -1:
		t.cnt[i] = 0
	case 2:
		t.cnt[i] = r - l + 1 - t.cnt[i]
	}

	if op != 2 {
		t.lazy[i] = op
		return
	}

	switch t.lazy[i] {
	case 0:
		t.lazy[i] = op
	case 1:
		t.lazy[i] = -1
	case -1:
		t.lazy[i] = 1
	default:
		// t.lazy[i] = 2
		t.lazy[i] = 0
	}
}

func (t *Tree) push(i int, l int, r int) {
	mid := (l + r) >> 1
	if t.lazy[i] != 0 && l < r {
		t.apply(i*2+1, l, mid, t.lazy[i])
		t.apply(i*2+2, mid+1, r, t.lazy[i])
		t.lazy[i] = 0
	}
}

// op = 1 or -1, 2
func (t *Tree) Update(L int, R int, op int) {

	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.apply(i, l, r, op)
			return
		}
		t.push(i, l, r)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}

	f(0, 0, t.sz-1, L, R)
}

func (t *Tree) FindFirstZero() int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		t.push(i, l, r)
		mid := (l + r) >> 1
		if t.cnt[i*2+1] < mid-l+1 {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, t.sz-1)
}
