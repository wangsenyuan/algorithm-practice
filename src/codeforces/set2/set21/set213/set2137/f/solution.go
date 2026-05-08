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
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	y := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &y[i])
	}
	return solve(x, y)
}

func solve(x []int, y []int) int {
	n := len(x)
	// x[i] <= 2 * n
	tr := NewSegTree(2*n + 1)

	var res int
	for r := range n {
		// 左边最大的l, x[l] >= y[r]
		if x[r] == y[r] {
			res += (r + 1) * (n - r)
		} else {
			l := tr.Get(max(x[r], y[r]), 2*n+1)
			res += (l + 1) * (n - r)
		}
		tr.Update(x[r], r)
	}

	return res
}

type SegTree []int

func NewSegTree(n int) SegTree {
	res := make(SegTree, 2*n)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	return res
}

func (t SegTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	if t[p] >= v {
		return
	}
	t[p] = v
	for p > 1 {
		t[p>>1] = max(t[p], v)
		p >>= 1
	}
}

func (t SegTree) Get(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	res := -1
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
