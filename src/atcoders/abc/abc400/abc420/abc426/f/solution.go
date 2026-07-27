package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r, k := cur[0]-1, cur[1]-1, cur[2]
		ans[i] = tr.buy(l, r, k)
	}

	return ans
}

type Tree struct {
	arr  []int // minimum remaining stock; sold-out products use inf
	cnt  []int // number of products not yet marked sold out
	lazy []int // pending addition to all stocks in the segment
}

func NewTree(a []int) *Tree {
	n := len(a)
	arr := make([]int, 4*n)
	cnt := make([]int, 4*n)
	lazy := make([]int, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			arr[i] = a[l]
			cnt[i]++
		} else {
			mid := (l + r) >> 1
			f(i*2+1, l, mid)
			f(i*2+2, mid+1, r)
			arr[i] = min(arr[i*2+1], arr[i*2+2])
			cnt[i] = cnt[i*2+1] + cnt[i*2+2]
		}
	}
	f(0, 0, n-1)
	return &Tree{arr: arr, cnt: cnt, lazy: lazy}
}

func (t *Tree) apply(i int, val int) {
	t.arr[i] -= val
	t.lazy[i] += val
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.apply(i*2+1, t.lazy[i])
		t.apply(i*2+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) buy(L int, R int, v int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if t.cnt[i] == 0 {
			return 0
		}
		if t.arr[i] > v && L <= l && r <= R {
			t.apply(i, v)
			return v * t.cnt[i]
		}
		var ans int
		if l == r {
			ans = t.arr[i]
			t.arr[i] = 1 << 60
			t.cnt[i] = 0
			return ans
		}
		t.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			ans += f(i*2+1, l, mid)
		}
		if mid < R {
			ans += f(i*2+2, mid+1, r)
		}
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
		t.arr[i] = min(t.arr[i*2+1], t.arr[i*2+2])
		return ans
	}
	return f(0, 0, len(t.arr)/4-1)
}
