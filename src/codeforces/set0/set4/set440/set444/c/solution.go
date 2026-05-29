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

	res := drive(reader)

	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	queries := make([][]int, m)
	for i := range m {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			queries[i] = make([]int, 4)
		} else {
			queries[i] = make([]int, 3)
		}
		queries[i][0] = t
		for j := 1; j < len(queries[i]); j++ {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {
	tr := NewTree(n + 1)
	var res []int
	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			tr.modify(l, r, x)
		} else {
			l, r := cur[1], cur[2]
			res = append(res, tr.query(l, r))
		}
	}
	return res
}

type Tree struct {
	mark []int
	sum  []int
	add  []int
	sz   int
}

func NewTree(n int) *Tree {
	mark := make([]int, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			mark[i] = l
		} else {
			mid := (l + r) >> 1
			f(i*2+1, l, mid)
			f(i*2+2, mid+1, r)
		}
	}
	f(0, 0, n-1)
	return &Tree{mark, make([]int, 4*n), make([]int, 4*n), n}
}

func (tr *Tree) clear(i int, l int, r int, x int) {
	if tr.mark[i] > 0 {
		tr.add[i] += abs(tr.mark[i] - x)
		tr.sum[i] += abs(tr.mark[i]-x) * (r - l + 1)
	} else {
		mid := (l + r) >> 1
		tr.clear(i*2+1, l, mid, x)
		tr.clear(i*2+2, mid+1, r, x)
		tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2] + tr.add[i]*(r-l+1)
	}
	tr.mark[i] = -1
}

func abs(num int) int {
	return max(num, -num)
}

func (tr *Tree) modify(L int, R int, x int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.clear(i, l, r, x)
			tr.mark[i] = x
			return
		}
		if tr.mark[i] > 0 {
			tr.mark[i*2+1] = tr.mark[i]
			tr.mark[i*2+2] = tr.mark[i]
			tr.mark[i] = 0
		}
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.mark[i] = 0
		tr.sum[i] = tr.sum[i*2+1] + tr.sum[i*2+2] + tr.add[i]*(r-l+1)
	}
	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) query(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.sum[i]
		}
		mid := (l + r) >> 1
		res := tr.add[i] * (R - L + 1)

		if L <= mid {
			res += f(i*2+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(L, mid+1), R)
		}

		return res
	}
	return f(0, 0, tr.sz-1, L, R)
}
