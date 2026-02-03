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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, m)
	for i := range m {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		queries[i] = []int{l, r}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	last := make(map[int]int)
	n := len(a)
	qs := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		qs[r] = append(qs[r], i)
	}

	ans := make([]int, len(queries))

	pos := NewSegTree(n)
	for i := range n {
		if j, ok := last[a[i]]; ok {
			pos.Update(j, i-j)
		}

		for _, j := range qs[i] {
			l := queries[j][0] - 1
			ans[j] = -1
			w := pos.Get(l, i+1)
			if w < inf {
				ans[j] = w
			}
		}

		last[a[i]] = i
	}

	return ans
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make(SegTree, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return SegTree(arr)
}

func (t SegTree) Update(p int, v int) {
	p += len(t) / 2
	t[p] = v
	for i := p; i > 1; i >>= 1 {
		t[i>>1] = min(t[i], t[i^1])
	}
}

func (t SegTree) Get(l int, r int) int {
	l += len(t) / 2
	r += len(t) / 2
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
